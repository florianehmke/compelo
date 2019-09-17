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

	s.createProject("My Project", "secret", 1)
	s.createProjectWithUniqueConstraintViolation("My Project")

	s.selectProject("My Project", "secret")
	s.selectProjectWithWrongPassword()

	s.createGame("My Game", 1)
	s.createGameWithUniqueConstraintViolation("My Game")

	s.createPlayer("Player 1", 1)
	s.createPlayer("Player 2", 2)
	s.createPlayerWithUniqueConstraintViolation("Player 2")

	s.createMatch(1)
	s.getMatchByID()
}

func (s *suite) createProject(name, pw string, expectedID uint) {
	b := gin.H{
		"name":     name,
		"password": pw,
	}
	w := s.requestWithBody("POST", "/api/create-project", b)
	response := &compelo.Project{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, expectedID, response.ID)
	assert.Equal(s.t, name, response.Name)
}

func (s *suite) createProjectWithUniqueConstraintViolation(name string) {
	b := gin.H{
		"name":     name,
		"password": "12345",
	}
	w := s.requestWithBody("POST", "/api/create-project", b)
	assert.Equal(s.t, http.StatusBadRequest, w.Code)
	assert.Contains(s.t, w.Body.String(), "UNIQUE constraint failed")
}

func (s *suite) selectProject(name, pw string) {
	type token struct {
		Code   int       `json:"code"`
		Expire time.Time `json:"expire"`
		Token  string    `json:"token"`
	}

	b := gin.H{
		"name":     name,
		"password": pw,
	}
	w := s.requestWithBody("POST", "/api/select-project", b)
	response := &token{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusOK, w.Code)
	assert.NotNil(s.t, response.Expire)
	assert.NotEmpty(s.t, response.Token)
	s.token = response.Token
}

func (s *suite) selectProjectWithWrongPassword() {
	b := gin.H{
		"name":     "foo",
		"password": "bar",
	}
	w := s.requestWithBody("POST", "/api/select-project", b)
	assert.Equal(s.t, http.StatusUnauthorized, w.Code)
	assert.Contains(s.t, w.Body.String(), "incorrect Username or Password")
}

func (s *suite) createGame(name string, expectedID uint) {
	b := gin.H{
		"name": name,
	}
	w := s.requestWithBody("POST", "/api/project/games", b)
	response := &compelo.Game{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, name, response.Name)
	assert.Equal(s.t, expectedID, response.ID)
}

func (s *suite) createGameWithUniqueConstraintViolation(name string) {
	b := gin.H{
		"name": name,
	}
	w := s.requestWithBody("POST", "/api/project/games", b)
	assert.Equal(s.t, http.StatusBadRequest, w.Code)
	assert.Contains(s.t, w.Body.String(), "UNIQUE constraint failed")
}

func (s *suite) createPlayer(name string, expectedID uint) {
	b := gin.H{
		"name": name,
	}
	w := s.requestWithBody("POST", "/api/project/players", b)
	response := &compelo.Player{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, name, response.Name)
	assert.Equal(s.t, expectedID, response.ID)
}

func (s *suite) createPlayerWithUniqueConstraintViolation(name string) {
	b := gin.H{
		"name": name,
	}
	w := s.requestWithBody("POST", "/api/project/players", b)
	assert.Equal(s.t, http.StatusBadRequest, w.Code)
	assert.Contains(s.t, w.Body.String(), "UNIQUE constraint failed")
}

func (s *suite) createMatch(expectedID uint) {
	b := gin.H{
		"teams": []gin.H{
			{
				"playerIds": []int{1},
				"score":     3,
				"winner":    false,
			},
			{
				"playerIds": []int{2},
				"score":     5,
				"winner":    true,
			},
		},
	}

	w := s.requestWithBody("POST", "/api/project/games/1/matches", b)
	response := &compelo.Match{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusCreated, w.Code)
	assert.Equal(s.t, expectedID, response.ID)
	assert.Equal(s.t, uint(1), response.GameID)
}

func (s *suite) getMatchByID() {
	w := s.request("GET", "/api/project/games/1/matches/1")
	response := &match.MatchData{}
	mustUnmarshal(s.t, w.Body.Bytes(), response)
	assert.Equal(s.t, http.StatusOK, w.Code)
	assert.Equal(s.t, uint(1), response.ID)
	assert.Equal(s.t, 2, len(response.Teams))
	assert.Equal(s.t, uint(1), response.GameID)

	// Team 1
	assert.True(s.t, response.Teams[0].Winner)
	assert.Equal(s.t, 5, response.Teams[0].Score)
	assert.Equal(s.t, uint(1), response.Teams[0].MatchID)
	assert.Equal(s.t, uint(2), response.Teams[0].ID)
	assert.Equal(s.t, 1, len(response.Teams[0].Players))
	assert.Equal(s.t, uint(2), response.Teams[0].Players[0].ID)

	// Team 2
	assert.False(s.t, response.Teams[1].Winner)
	assert.Equal(s.t, 3, response.Teams[1].Score)
	assert.Equal(s.t, uint(1), response.Teams[1].MatchID)
	assert.Equal(s.t, uint(1), response.Teams[1].ID)
	assert.Equal(s.t, 1, len(response.Teams[1].Players))
	assert.Equal(s.t, uint(1), response.Teams[1].Players[0].ID)
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
