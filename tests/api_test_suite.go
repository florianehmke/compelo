package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"compelo/api/handler"
	"compelo/api/router"
	"compelo/api/security"
	"compelo/command"
	"compelo/event"
	"compelo/query"
)

type apiTestSuite struct {
	testing *testing.T
	dbName  string

	q *query.Service
	c *command.Service

	handler http.Handler
	token   string

	// 1. Project
	projectRequest handler.CreateProjectRequest
	projectGUID    string
	project        query.Project

	// 2. Game
	gameRequest handler.CreateGameRequest
	gameGUID    string
	game        query.Game

	// 3. Players
	playerRequests map[int]handler.CreatePlayerRequest
	playerGUIDs    map[int]string
	players        map[int]query.Player

	// 4. Matches
	matchRequests map[int]handler.CreateMatchRequest
	matchGUIDs    map[int]string
	matches       map[int]query.Match
	gameStats     query.GameStats
	playerStats   query.PlayerStats

	expectedMatches     map[int]query.Match
	expectedPlayerStats map[int]query.PlayerStats
	expectedGameStats   map[int]query.GameStats

	// 5. Competitions
	competitionRequest handler.CreateCompetitionRequest
	competitionGUID    string
}

func createAPITestSuite(t *testing.T) *apiTestSuite {
	dbName := uuid.New().String() + ".db"

	// Create event store.
	bus := event.NewBus()
	store := event.NewStore(bus, dbName)

	// Setup query.
	q := query.NewService(bus)

	// Load all events from db (rehydrates queries).
	events, err := store.LoadEvents()
	assert.Nil(t, err)

	// Setup command (from existing events).
	c := command.NewService(store, events)

	hdl := handler.New(q, c)
	sec := security.New(q, 60, "test")
	mux := router.New(hdl, sec)

	return &apiTestSuite{
		dbName:      dbName,
		testing:     t,
		handler:     mux,
		c:           c,
		q:           q,
		playerGUIDs: make(map[int]string),
		players:     make(map[int]query.Player),
		matchGUIDs:  make(map[int]string),
		matches:     make(map[int]query.Match),
	}
}

func (s *apiTestSuite) createProject() {
	w := s.request("POST", "/api/projects", s.projectRequest)

	response := &command.Response{}
	s.mustUnmarshal(w.Body.Bytes(), response)

	s.assertEqual(http.StatusCreated, w.Code)
	s.assertNotEmpty(response.GUID)
	s.projectGUID = response.GUID
}

func (s *apiTestSuite) listProjects() {
	w := s.request("GET", "/api/projects", nil)

	var response []query.Project
	s.mustUnmarshal(w.Body.Bytes(), &response)

	s.assertEqual(http.StatusOK, w.Code)
	s.assertTrue(len(response) == 1)
	s.assertEqual(response[0].Name, s.projectRequest.Name)
	s.assertEqual(response[0].GUID, s.projectGUID)
	s.project = response[0]
}

func (s *apiTestSuite) selectProject() {
	w := s.request("POST", "/api/login", security.AuthRequest{
		ProjectGUID: s.projectGUID,
		Password:    s.projectRequest.Password,
	})

	type token struct {
		Code   int       `json:"code"`
		Expire time.Time `json:"expire"`
		Token  string    `json:"token"`
	}

	response := &token{}
	s.mustUnmarshal(w.Body.Bytes(), response)

	s.assertEqual(http.StatusOK, w.Code)
	s.assertNotNil(response.Expire)
	s.assertNotEmpty(response.Token)
	s.token = response.Token
}

func (s *apiTestSuite) createGame() {
	w := s.request("POST", "/api/projects/"+s.projectGUID+"/games", s.gameRequest)

	response := &command.Response{}
	s.mustUnmarshal(w.Body.Bytes(), response)

	s.assertEqual(http.StatusCreated, w.Code)
	s.assertNotEmpty(response.GUID)
	s.gameGUID = response.GUID
}

func (s *apiTestSuite) listGames() {
	w := s.request("GET", "/api/projects/"+s.projectGUID+"/games", nil)

	var response []query.Game
	s.mustUnmarshal(w.Body.Bytes(), &response)

	s.assertEqual(http.StatusOK, w.Code)
	s.assertTrue(len(response) == 1)
	s.assertEqual(response[0].Name, s.gameRequest.Name)
	s.assertEqual(response[0].GUID, s.gameGUID)
	s.game = response[0]
}

func (s *apiTestSuite) createPlayers() {
	for i, r := range s.playerRequests {
		w := s.request("POST", "/api/projects/"+s.projectGUID+"/players", r)

		var response command.Response
		s.mustUnmarshal(w.Body.Bytes(), &response)

		s.assertEqual(http.StatusCreated, w.Code)
		s.assertNotEmpty(response.GUID)
		s.playerGUIDs[i] = response.GUID
	}
}

func (s *apiTestSuite) listPlayers() {
	w := s.request("GET", "/api/projects/"+s.projectGUID+"/players", nil)

	var response []query.Player
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertEqual(http.StatusOK, w.Code)

	s.assertTrue(len(response) == len(s.playerRequests))
	for i, player := range response {
		s.players[i] = player
		s.assertEqual(player.GUID, s.playerGUIDs[i])
		s.assertEqual(player.Name, s.playerRequests[i].Name)
	}
}

func (s *apiTestSuite) createMatches() {
	for i, r := range s.matchRequests {
		w := s.request("POST", "/api/projects/"+s.projectGUID+"/games/"+s.gameGUID+"/matches", r)

		response := &command.Response{}
		s.mustUnmarshal(w.Body.Bytes(), response)

		s.assertEqual(http.StatusCreated, w.Code)
		s.assertNotEmpty(response.GUID)
		s.matchGUIDs[i] = response.GUID
	}
}

func (s *apiTestSuite) listMatches() {
	w := s.request("GET", "/api/projects/"+s.projectGUID+"/games/"+s.gameGUID+"/matches", nil)

	var response []query.Match
	s.mustUnmarshal(w.Body.Bytes(), &response)

	s.assertEqual(http.StatusOK, w.Code)
	s.assertTrue(len(response) == len(s.matchRequests))

	for _, match := range response {
		matchIndex := -1

		// Map iteration is not ordered, the matches are created in a random order.
		// Therefore we have to determine the index of the match that is checked here
		// by its guid.
		for i, guid := range s.matchGUIDs {
			if guid == match.GUID {
				matchIndex = i
				break
			}
		}

		s.assertEqual(s.matchGUIDs[matchIndex], match.GUID)
		s.assertEqual(s.gameGUID, match.GameGUID)
		s.assertEqual(s.projectGUID, match.ProjectGUID)

		matchRequest := s.matchRequests[matchIndex]
		s.assertEqual(len(matchRequest.Teams), len(match.Teams))
		s.assertNotEmpty(match.CreatedDate)
		s.assertNotEmpty(match.UpdatedDate)
		s.matches[matchIndex] = match

		for j, team := range match.Teams {
			s.assertEqual(len(matchRequest.Teams[j].PlayerGUIDs), len(team.Players))
			s.assertEqual(matchRequest.Teams[j].Score, team.Score)
			s.assertNotEmpty(team.Result)
		}

		if s.expectedMatches != nil {
			expectedResponse := s.expectedMatches[matchIndex]
			s.assertEqual(expectedResponse.GUID, s.matches[matchIndex].GUID)
		}
	}
}

func (s *apiTestSuite) loadPlayerStats() {
	w := s.request("GET", "/api/projects/"+s.projectGUID+"/games/"+s.gameGUID+"/player-stats", nil)

	var response []query.PlayerStats
	s.mustUnmarshal(w.Body.Bytes(), &response)

	s.assertEqual(http.StatusOK, w.Code)
	s.assertTrue(len(response) == len(s.playerRequests))
}

func (s *apiTestSuite) loadGameStats() {
	w := s.request("GET", "/api/projects/"+s.projectGUID+"/games/"+s.gameGUID+"/game-stats", nil)

	var response query.GameStats
	s.mustUnmarshal(w.Body.Bytes(), &response)

	s.assertEqual(http.StatusOK, w.Code)
	s.assertTrue(len(response.MaxScoreDiff) == len(s.matchRequests))
	s.assertTrue(len(response.MaxScoreSum) == len(s.matchRequests))
}

func (s *apiTestSuite) createCompetition() {
	w := s.request("POST", "/api/projects/"+s.projectGUID+"/games/"+s.gameGUID+"/competitions", s.competitionRequest)

	response := &command.Response{}
	s.mustUnmarshal(w.Body.Bytes(), response)

	s.assertEqual(http.StatusCreated, w.Code)
	s.assertNotEmpty(response.GUID)
	s.competitionGUID = response.GUID
}

// ------ Helpers ------

func (s *apiTestSuite) request(method, path string, body interface{}) *httptest.ResponseRecorder {
	var req *http.Request
	var err error

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			s.testing.Error(err)
		}
		req, err = http.NewRequest(method, path, bytes.NewBuffer(b))
		if err != nil {
			s.testing.Error(err)
		}
	} else {
		req, err = http.NewRequest(method, path, nil)
		if err != nil {
			s.testing.Error(err)
		}
	}

	req.Header.Set("Content-Type", "application/json")
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}
	w := httptest.NewRecorder()
	s.handler.ServeHTTP(w, req)
	return w
}

func (s *apiTestSuite) mustUnmarshal(bytes []byte, target interface{}) {
	err := json.Unmarshal(bytes, target)
	if err != nil {
		s.testing.Error(err)
	}
}

func (s *apiTestSuite) assertEqual(expected, actual interface{}) {
	assert.Equal(s.testing, expected, actual)
}

func (s *apiTestSuite) assertTrue(value bool) {
	assert.True(s.testing, value)
}

func (s *apiTestSuite) assertNotNil(obj interface{}) {
	assert.NotNil(s.testing, obj)
}

func (s *apiTestSuite) assertNotEmpty(obj interface{}) {
	assert.NotEmpty(s.testing, obj)
}
