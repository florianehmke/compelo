package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	match   testMatch

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
		match: testMatch{
			teams: []testTeam{
				{
					players: []int{1, 2},
					score:   2,
				},
				{
					players: []int{3, 4},
					score:   1,
				},
			},
		},
		stats: map[string]testStats{
			"Hans": {
				rating:       1516,
				peakRating:   1516,
				lowestRating: 1500,
				gameCount:    1,
			},
			"Peter": {
				rating:       1516,
				peakRating:   1516,
				lowestRating: 1500,
				gameCount:    1,
			},
			"Kevin": {
				rating:       1484,
				peakRating:   1500,
				lowestRating: 1484,
				gameCount:    1,
			},
			"Arnold": {
				rating:       1484,
				peakRating:   1500,
				lowestRating: 1484,
				gameCount:    1,
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

	ts.createMatch()
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
		"name":     s.testData.projectName,
		"password": s.testData.projectPW,
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

func (s *testSuite) createMatch() {
	var teams []gin.H
	for _, t := range s.testData.match.teams {
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

func (s *testSuite) listMatches() {
	gameID := strconv.Itoa(int(s.testData.gameID))
	w := s.request("GET", "/api/project/games/"+gameID+"/matches")

	var response []match.MatchData
	s.assertEqual(http.StatusOK, w.Code)
	s.mustUnmarshal(w.Body.Bytes(), &response)
	s.assertTrue(len(response) == 1)
	s.assertTrue(len(response[0].Teams) == 2)

	for i, t := range response[0].Teams {
		s.assertTrue(t.Score == s.testData.match.teams[i].score)
		s.assertTrue(len(t.Players) == len(s.testData.match.teams[i].players))
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
