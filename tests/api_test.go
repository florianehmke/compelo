package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"compelo/api/handler"
	"compelo/api/router"
	"compelo/api/security"
	"compelo/command"
	"compelo/event"
	"compelo/query"
)

type JSON map[string]interface{}

type testSuite struct {
	testing  *testing.T
	handler  http.Handler
	testData testData
	token    string
}

func newTestSuite(t *testing.T, handler http.Handler, ts testData) *testSuite {
	return &testSuite{
		testing:  t,
		handler:  handler,
		testData: ts,
	}
}

type testData struct {
	projectName string
	projectPW   string
	projectGUID string // filled during test

	gameName string
	gameGUID string // filled during test

	players []testPlayer
	matches []testMatch

	stats map[string]testStats
}

type testPlayer struct {
	name string
	guid string
}

type testMatch struct {
	teams []testTeam // sort by score for assertions
}

type testTeam struct {
	players []int
	score   int
	result  query.Result
}

type testStats struct {
	rating       int
	peakRating   int
	lowestRating int
	gameCount    int
	winCount     int
	drawCount    int
	lossCount    int
}

func TestAPI(t *testing.T) {
	testData := testData{
		projectName: "Avengers",
		projectPW:   "secret",
		gameName:    "Fifa",
		players: []testPlayer{
			{name: "Hans"},
			{name: "Peter"},
			{name: "Kevin"},
			{name: "Arnold"},
		},
		matches: []testMatch{
			{
				teams: []testTeam{
					{
						players: []int{0, 1},
						score:   1,
						result:  query.Draw,
					},
					{
						players: []int{2, 3},
						score:   1,
						result:  query.Draw,
					},
				},
			},
			{
				teams: []testTeam{
					{
						players: []int{0, 1},
						score:   4,
						result:  query.Win,
					},
					{
						players: []int{2, 3},
						score:   1,
						result:  query.Loss,
					},
				},
			},
			{
				teams: []testTeam{
					{
						players: []int{0},
						score:   4,
						result:  query.Win,
					},
					{
						players: []int{1},
						score:   4,
						result:  query.Win,
					},
					{
						players: []int{2},
						score:   2,
						result:  query.Loss,
					},
				},
			},
		},
		stats: map[string]testStats{
			"Hans": {
				rating:       1523,
				peakRating:   1523,
				lowestRating: 1500,
				gameCount:    3,
				winCount:     2,
				drawCount:    1,
				lossCount:    0,
			},
			"Peter": {
				rating:       1523,
				peakRating:   1523,
				lowestRating: 1500,
				gameCount:    3,
				winCount:     2,
				drawCount:    1,
				lossCount:    0,
			},
			"Kevin": {
				rating:       1470,
				peakRating:   1500,
				lowestRating: 1470,
				gameCount:    3,
				winCount:     0,
				drawCount:    1,
				lossCount:    2,
			},
			"Arnold": {
				rating:       1484,
				peakRating:   1500,
				lowestRating: 1484,
				gameCount:    2,
				winCount:     0,
				drawCount:    1,
				lossCount:    1,
			},
		},
	}

	defer os.Remove("api_test.db")

	// Create event store.
	bus := event.NewBus()
	store := event.NewStore(bus, "api_test.db")

	// Setup query.
	query := query.New(bus)

	// Load all events from db (rehydrates queries).
	events, err := store.LoadEvents()
	assert.Nil(t, err)

	// Setup command (from existing events).
	command := command.New(store, events)

	hdl := handler.New(query, command)
	sec := security.New(query, 60, "test")
	mux := router.New(hdl, sec)

	ts := newTestSuite(t, mux, testData)

	ts.createProject()
	ts.listProjects()
	ts.selectProject()

	ts.createPlayers()
	ts.listPlayers()

	ts.createGame()
	ts.listGames()

	ts.createMatches()
	ts.listMatches()

	ts.loadPlayerStats()
	ts.loadGameStats()
}

func (s *testSuite) createProject() {
	b := JSON{
		"name":     s.testData.projectName,
		"password": s.testData.projectPW,
	}
	w := s.requestWithBody("POST", "/api/projects", b)

	response := &query.Project{}
	s.assertEqual(http.StatusCreated, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), response)
	s.assertTrue(response.GUID != "")
}

func (s *testSuite) listProjects() {
	w := s.request("GET", "/api/projects")

	var response []query.Project
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == 1)
	s.assertEqual(response[0].Name, s.testData.projectName)
	s.testData.projectGUID = response[0].GUID
}

func (s *testSuite) selectProject() {
	b := JSON{
		"projectGuid": s.testData.projectGUID,
		"password":    s.testData.projectPW,
	}
	w := s.requestWithBody("POST", "/api/login", b)

	type token struct {
		Code   int       `json:"code"`
		Expire time.Time `json:"expire"`
		Token  string    `json:"token"`
	}
	response := &token{}
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), response)
	s.assertNotNil(response.Expire)
	s.assertNotEmpty(response.Token)
	s.token = response.Token
}

func (s *testSuite) createPlayers() {
	for i, p := range s.testData.players {
		b := JSON{
			"name": p.name,
		}
		w := s.requestWithBody("POST", "/api/projects/"+s.testData.projectGUID+"/players", b)
		var response command.Response
		s.mustUnmarshal(w.Body.Bytes(), &response)
		s.assertEqual(http.StatusCreated, w.Code)
		s.assertTrue(response.GUID != "")
		s.testData.players[i].guid = response.GUID
	}
}

func (s *testSuite) listPlayers() {
	w := s.request("GET", "/api/projects/"+s.testData.projectGUID+"/players")

	var response []query.Player
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == len(s.testData.players))
	for i, r := range response {
		s.assertEqual(r.Name, s.testData.players[i].name)
	}
}

func (s *testSuite) createGame() {
	b := JSON{
		"name": s.testData.gameName,
	}
	w := s.requestWithBody("POST", "/api/projects/"+s.testData.projectGUID+"/games", b)
	s.assertEqual(http.StatusCreated, w.Code)
}

func (s *testSuite) listGames() {
	w := s.request("GET", "/api/projects/"+s.testData.projectGUID+"/games")

	var response []query.Game
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == 1)
	s.assertEqual(response[0].Name, s.testData.gameName)
	s.testData.gameGUID = response[0].GUID
}

func (s *testSuite) createMatches() {
	for _, m := range s.testData.matches {
		var teams []JSON
		for _, t := range m.teams {
			var playerGUIDs []string
			for _, idx := range t.players {
				playerGUIDs = append(playerGUIDs, s.testData.players[idx].guid)
			}
			teams = append(teams, JSON{
				"playerGuids": playerGUIDs,
				"score":       t.score,
			})
		}
		body := JSON{"teams": teams}

		gameGUID := s.testData.gameGUID
		w := s.requestWithBody("POST", "/api/projects/"+s.testData.projectGUID+"/games/"+gameGUID+"/matches", body)

		response := &command.Response{}
		s.assertEqual(http.StatusCreated, w.Code)
		s.mustUnmarshal(w.Body.Bytes(), response)
		s.assertTrue(response.GUID != "")
	}
}

func (s *testSuite) listMatches() {
	gameGUID := s.testData.gameGUID
	w := s.request("GET", "/api/projects/"+s.testData.projectGUID+"/games/"+gameGUID+"/matches")

	var response []query.Match
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == len(s.testData.matches))

	matchCount := len(s.testData.matches)
	for i, expectedMatch := range s.testData.matches {
		actualMatch := response[matchCount-i-1] // matches are sorted: newest first
		s.assertEqual(len(expectedMatch.teams), len(actualMatch.Teams))

		for j, expectedTeam := range expectedMatch.teams {
			actualTeam := actualMatch.Teams[j]
			s.assertEqual(len(expectedTeam.players), len(actualTeam.Players))
			s.assertEqual(expectedTeam.score, actualTeam.Score)
			s.assertEqual(expectedTeam.result, actualTeam.Result)
		}
	}
}

func (s *testSuite) loadPlayerStats() {
	gameGUID := s.testData.gameGUID
	w := s.request("GET", "/api/projects/"+s.testData.projectGUID+"/games/"+gameGUID+"/player-stats")

	var response []query.PlayerStats
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == len(s.testData.players))

	for _, v := range response {
		s.assertEqual(s.testData.stats[v.Name].rating, v.Current.Rating)
		s.assertEqual(s.testData.stats[v.Name].peakRating, v.Current.PeakRating)
		s.assertEqual(s.testData.stats[v.Name].lowestRating, v.Current.LowestRating)
		s.assertEqual(s.testData.stats[v.Name].gameCount, v.Current.GameCount)
		s.assertEqual(s.testData.stats[v.Name].winCount, v.Current.WinCount)
		s.assertEqual(s.testData.stats[v.Name].drawCount, v.Current.DrawCount)
		s.assertEqual(s.testData.stats[v.Name].lossCount, v.Current.LossCount)
	}
}

func (s *testSuite) loadGameStats() {
	gameGUID := s.testData.gameGUID
	w := s.request("GET", "/api/projects/"+s.testData.projectGUID+"/games/"+gameGUID+"/game-stats")

	var response query.GameStats
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)

	// Only three test games so far, all should appear in stats.
	s.assertTrue(len(response.MaxScoreDiff) == len(s.testData.matches))
	s.assertTrue(len(response.MaxScoreSum) == len(s.testData.matches))

	// The third test-game, 1v1v1 (4:4:2)
	s.assertTrue(len(response.MaxScoreSum[0].Teams) == 3)

	// The second test-game, 1v1 (4:1)
	s.assertTrue(response.MaxScoreDiff[0].Teams[0].Score == 4)
	s.assertTrue(response.MaxScoreDiff[0].Teams[1].Score == 1)
}

// ------ Helpers ------

func (s *testSuite) requestWithBody(method, path string, body JSON) *httptest.ResponseRecorder {
	b, err := json.Marshal(body)
	if err != nil {
		s.testing.Error(err)
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(b))
	if err != nil {
		s.testing.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}
	w := httptest.NewRecorder()
	s.handler.ServeHTTP(w, req)
	return w
}

func (s *testSuite) request(method, path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		s.testing.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}
	w := httptest.NewRecorder()
	s.handler.ServeHTTP(w, req)
	return w
}

func (s *testSuite) mustUnmarshal(bytes []byte, target interface{}) {
	err := json.Unmarshal(bytes, target)
	log.Println(string(bytes))
	if err != nil {
		s.testing.Error(err)
	}
}

func (s *testSuite) assertEqual(expected, actual interface{}) {
	assert.Equal(s.testing, expected, actual)
}

func (s *testSuite) assertTrue(value bool) {
	assert.True(s.testing, value)
}

func (s *testSuite) assertNotNil(obj interface{}) {
	assert.NotNil(s.testing, obj)
}

func (s *testSuite) assertNotEmpty(obj interface{}) {
	assert.NotEmpty(s.testing, obj)
}
