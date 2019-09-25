package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"compelo/api"
	"compelo/game"
	"compelo/match"
	"compelo/player"
	"compelo/project"
	"compelo/stats"
)

type testSuite struct {
	testing  *testing.T
	router   *gin.Engine
	testData testData
	token    string
}

func newTestSuite(t *testing.T, r *gin.Engine, ts testData) *testSuite {
	return &testSuite{
		testing:  t,
		router:   r,
		testData: ts,
	}
}

type testData struct {
	projectName string
	projectPW   string
	projectID   uint // filled during test

	gameName string
	gameID   uint // filled during test

	players []testPlayer
	matches []testMatch

	stats map[string]testStats
}

type testPlayer struct {
	name string
}

type testMatch struct {
	teams []testTeam // sort by score for assertions
}

type testTeam struct {
	players []int
	score   int
	result  string
}

type testStats struct {
	rating       int
	peakRating   int
	lowestRating int
	gameCount    int
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
						players: []int{1, 2},
						score:   1,
						result:  "Draw",
					},
					{
						players: []int{3, 4},
						score:   1,
						result:  "Draw",
					},
				},
			},
			{
				teams: []testTeam{
					{
						players: []int{1, 2},
						score:   2,
						result:  "Win",
					},
					{
						players: []int{3, 4},
						score:   1,
						result:  "Loss",
					},
				},
			},
			{
				teams: []testTeam{
					{
						players: []int{1},
						score:   2,
						result:  "Win",
					},
					{
						players: []int{2},
						score:   2,
						result:  "Win",
					},
					{
						players: []int{3},
						score:   1,
						result:  "Loss",
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
			},
			"Peter": {
				rating:       1523,
				peakRating:   1523,
				lowestRating: 1500,
				gameCount:    3,
			},
			"Kevin": {
				rating:       1470,
				peakRating:   1500,
				lowestRating: 1470,
				gameCount:    3,
			},
			"Arnold": {
				rating:       1484,
				peakRating:   1500,
				lowestRating: 1484,
				gameCount:    2,
			},
		},
	}

	ts := newTestSuite(t, api.Setup("file::memory:", "test", false), testData)

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
}

func (s *testSuite) createProject() {
	b := gin.H{
		"name":     s.testData.projectName,
		"password": s.testData.projectPW,
	}
	w := s.requestWithBody("POST", "/api/create-project", b)

	response := &project.Project{}
	s.assertEqual(http.StatusCreated, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), response)
	s.assertTrue(response.ID > 0)
	s.assertEqual(s.testData.projectName, response.Name)
}

func (s *testSuite) listProjects() {
	w := s.request("GET", "/api/projects")

	var response []project.Project
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == 1)
	s.assertEqual(response[0].Name, s.testData.projectName)
	s.testData.projectID = response[0].ID
}

func (s *testSuite) selectProject() {
	b := gin.H{
		"projectName": s.testData.projectName,
		"password":    s.testData.projectPW,
	}
	w := s.requestWithBody("POST", "/api/select-project", b)

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
	for _, p := range s.testData.players {
		b := gin.H{
			"name": p.name,
		}
		w := s.requestWithBody("POST", "/api/project/players", b)
		s.assertEqual(http.StatusCreated, w.Code)
	}
}

func (s *testSuite) listPlayers() {
	w := s.request("GET", "/api/project/players")

	var response []player.Player
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == len(s.testData.players))
	for i, r := range response {
		s.assertEqual(r.Name, s.testData.players[i].name)
	}
}

func (s *testSuite) createGame() {
	b := gin.H{
		"name": s.testData.gameName,
	}
	w := s.requestWithBody("POST", "/api/project/games", b)
	s.assertEqual(http.StatusCreated, w.Code)
}

func (s *testSuite) listGames() {
	w := s.request("GET", "/api/project/games")

	var response []game.Game
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == 1)
	s.assertEqual(response[0].Name, s.testData.gameName)
	s.testData.gameID = response[0].ID
}

func (s *testSuite) createMatches() {
	for _, m := range s.testData.matches {
		var teams []gin.H
		for _, t := range m.teams {
			teams = append(teams, gin.H{
				"playerIds": t.players,
				"score":     t.score,
			})
		}
		body := gin.H{"teams": teams}

		gameID := strconv.Itoa(int(s.testData.gameID))
		w := s.requestWithBody("POST", "/api/project/games/"+gameID+"/matches", body)

		response := &match.Match{}
		s.assertEqual(http.StatusCreated, w.Code)
		s.mustUnmarshal(w.Body.Bytes(), response)
		s.assertTrue(response.ID > 0)
		s.assertEqual(response.GameID, s.testData.gameID)
	}
}

func (s *testSuite) listMatches() {
	gameID := strconv.Itoa(int(s.testData.gameID))
	w := s.request("GET", "/api/project/games/"+gameID+"/matches")

	var response []match.MatchData
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == len(s.testData.matches))

	// Sort results by ID so that they are
	// in the same order as the test data.
	sort.Slice(response, func(i, j int) bool {
		return response[i].ID < (response[j].ID)
	})

	for i, expectedMatch := range s.testData.matches {
		actualMatch := response[i]
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
	gameID := strconv.Itoa(int(s.testData.gameID))
	w := s.request("GET", "/api/project/games/"+gameID+"/players")

	var response []stats.Player
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == len(s.testData.players))

	for _, v := range response {
		s.assertEqual(s.testData.stats[v.Name].rating, v.Rating)
		s.assertEqual(s.testData.stats[v.Name].peakRating, v.PeakRating)
		s.assertEqual(s.testData.stats[v.Name].lowestRating, v.LowestRating)
		s.assertEqual(s.testData.stats[v.Name].gameCount, v.GameCount)
	}
}

// ------ Helpers ------

func (s *testSuite) requestWithBody(method, path string, body gin.H) *httptest.ResponseRecorder {
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
	s.router.ServeHTTP(w, req)
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
	s.router.ServeHTTP(w, req)
	return w
}

func (s *testSuite) mustUnmarshal(bytes []byte, target interface{}) {
	err := json.Unmarshal(bytes, target)
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
