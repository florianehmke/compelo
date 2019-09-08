package compelo_test

import (
	"bytes"
	"compelo"
	"compelo/match"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"compelo/api"
)

type suite struct {
	t     *testing.T
	r     *gin.Engine
	token string
}

func Test(t *testing.T) {
	s := suite{t: t, r: api.Setup("file::memory:")}

	s.createProject()
	s.selectProject()
	s.createGame()
	s.createPlayer("Player 1", 1)
	s.createPlayer("Player 2", 2)
	s.createMatch()
	s.getMatchByID()
}

func (s *suite) createProject() {
	b := gin.H{
		"name":     "My Project",
		"password": "secret",
	}
	w := s.requestWithBody("POST", "/create-project", b)
	response := &compelo.Project{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, uint(1), response.ID)
	assert.Equal(s.t, "My Project", response.Name)
}

func (s *suite) selectProject() {
	type token struct {
		Code   int       `json:"code"`
		Expire time.Time `json:"expire"`
		Token  string    `json:"token"`
	}

	b := gin.H{
		"name":     "My Project",
		"password": "secret",
	}
	w := s.requestWithBody("POST", "/select-project", b)
	response := &token{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusOK, w.Code)
	assert.Equal(s.t, 200, response.Code)
	assert.NotNil(s.t, response.Expire)
	assert.NotEmpty(s.t, response.Token)
	s.token = response.Token
}

func (s *suite) createGame() {
	b := gin.H{
		"name": "My Game",
	}
	w := s.requestWithBody("POST", "/project/games", b)
	response := &compelo.Game{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, "My Game", response.Name)
	assert.Equal(s.t, uint(1), response.ID)
}

func (s *suite) createPlayer(name string, expectedID uint) {
	b := gin.H{
		"name": name,
	}
	w := s.requestWithBody("POST", "/project/players", b)
	response := &compelo.Player{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, name, response.Name)
	assert.Equal(s.t, expectedID, response.ID)
}

func (s *suite) createMatch() {
	b := gin.H{
		"name":   "My Player",
		"gameId": 1,
		"playerTeamMap": gin.H{
			"1": 1,
			"2": 2,
		},
		"teamScoreMap": gin.H{
			"1": 3,
			"2": 5,
		},
		"teams":             2,
		"winnerMatchTeamId": 2,
	}
	w := s.requestWithBody("POST", "/project/games/1/matches", b)
	response := &compelo.Match{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, uint(1), response.ID)
	assert.Equal(s.t, uint(1), response.GameID)
	assert.Equal(s.t, uint(2), response.WinnerMatchTeamID)
}

func (s *suite) getMatchByID() {
	w := s.request("GET", "/project/games/1/matches/1")
	response := &match.CompleteMatch{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusOK, w.Code)
	assert.Equal(s.t, uint(1), response.Match.ID)
	assert.Equal(s.t, 2, len(response.MatchPlayers))
	assert.Equal(s.t, 2, len(response.MatchTeams))
	assert.Equal(s.t, uint(1), response.Match.GameID)
	assert.Equal(s.t, uint(2), response.Match.WinnerMatchTeamID)

	// Team 1
	assert.Equal(s.t, 3, response.MatchTeams[0].Score)
	assert.Equal(s.t, uint(1), response.MatchTeams[0].MatchID)
	assert.Equal(s.t, uint(1), response.MatchTeams[0].ID)

	// Team 2
	assert.Equal(s.t, 5, response.MatchTeams[1].Score)
	assert.Equal(s.t, uint(1), response.MatchTeams[1].MatchID)
	assert.Equal(s.t, uint(2), response.MatchTeams[1].ID)
}

func (s *suite) requestWithBody(method, path string, body gin.H) *httptest.ResponseRecorder {
	b, err := json.Marshal(body)
	if err != nil {
		s.t.Error(err)
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(b))
	if err != nil {
		s.t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	return w
}

func (s *suite) request(method, path string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		s.t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if s.token != "" {
		req.Header.Set("Authorization", "Bearer "+s.token)
	}
	w := httptest.NewRecorder()
	s.r.ServeHTTP(w, req)
	return w
}

func mustUnmarshal(t *testing.T, bytes []byte, target interface{}) {
	err := json.Unmarshal(bytes, target)
	if err != nil {
		t.Error(err)
	}
}
