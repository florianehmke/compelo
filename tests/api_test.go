package tests

// type JSON map[string]interface{}

// type testSuite struct {
// 	testing  *testing.T
// 	handler  http.Handler
// 	testData testData
// 	token    string
// }

// func newTestSuite(t *testing.T, handler http.Handler, ts testData) *testSuite {
// 	return &testSuite{
// 		testing:  t,
// 		handler:  handler,
// 		testData: ts,
// 	}
// }

// type testData struct {
// 	projectName     string
// 	projectPW       string
// 	projectID       uint   // filled during test
// 	projectIDString string // filled during test

// 	gameName string
// 	gameID   uint // filled during test

// 	players []testPlayer
// 	matches []testMatch

// 	stats map[string]testStats
// }

// type testPlayer struct {
// 	name string
// }

// type testMatch struct {
// 	teams []testTeam // sort by score for assertions
// }

// type testTeam struct {
// 	players []int
// 	score   int
// 	result  db.Result
// }

// type testStats struct {
// 	rating       int
// 	peakRating   int
// 	lowestRating int
// 	gameCount    int
// 	winCount     int
// 	drawCount    int
// 	lossCount    int
// }

// func TestAPI(t *testing.T) {
// 	testData := testData{
// 		projectName: "Avengers",
// 		projectPW:   "secret",
// 		gameName:    "Fifa",
// 		players: []testPlayer{
// 			{name: "Hans"},
// 			{name: "Peter"},
// 			{name: "Kevin"},
// 			{name: "Arnold"},
// 		},
// 		matches: []testMatch{
// 			{
// 				teams: []testTeam{
// 					{
// 						players: []int{1, 2},
// 						score:   1,
// 						result:  db.Draw,
// 					},
// 					{
// 						players: []int{3, 4},
// 						score:   1,
// 						result:  db.Draw,
// 					},
// 				},
// 			},
// 			{
// 				teams: []testTeam{
// 					{
// 						players: []int{1, 2},
// 						score:   4,
// 						result:  db.Win,
// 					},
// 					{
// 						players: []int{3, 4},
// 						score:   1,
// 						result:  db.Loss,
// 					},
// 				},
// 			},
// 			{
// 				teams: []testTeam{
// 					{
// 						players: []int{1},
// 						score:   4,
// 						result:  db.Win,
// 					},
// 					{
// 						players: []int{2},
// 						score:   4,
// 						result:  db.Win,
// 					},
// 					{
// 						players: []int{3},
// 						score:   2,
// 						result:  db.Loss,
// 					},
// 				},
// 			},
// 		},
// 		stats: map[string]testStats{
// 			"Hans": {
// 				rating:       1523,
// 				peakRating:   1523,
// 				lowestRating: 1500,
// 				gameCount:    3,
// 				winCount:     2,
// 				drawCount:    1,
// 				lossCount:    0,
// 			},
// 			"Peter": {
// 				rating:       1523,
// 				peakRating:   1523,
// 				lowestRating: 1500,
// 				gameCount:    3,
// 				winCount:     2,
// 				drawCount:    1,
// 				lossCount:    0,
// 			},
// 			"Kevin": {
// 				rating:       1470,
// 				peakRating:   1500,
// 				lowestRating: 1470,
// 				gameCount:    3,
// 				winCount:     0,
// 				drawCount:    1,
// 				lossCount:    2,
// 			},
// 			"Arnold": {
// 				rating:       1484,
// 				peakRating:   1500,
// 				lowestRating: 1484,
// 				gameCount:    2,
// 				winCount:     0,
// 				drawCount:    1,
// 				lossCount:    1,
// 			},
// 		},
// 	}

// 	svc := compelo.NewService("file::memory:")
// 	hdl := handler.New(svc)
// 	sec := security.New(svc, 60, "test")
// 	mux := router.New(hdl, sec)

// 	ts := newTestSuite(t, mux, testData)

// 	ts.createProject()
// 	ts.listProjects()
// 	ts.selectProject()

// 	ts.createPlayers()
// 	ts.listPlayers()

// 	ts.createGame()
// 	ts.listGames()

// 	ts.createMatches()
// 	ts.listMatches()

// 	ts.loadPlayerStats()
// 	ts.loadGameStats()
// }

// func (s *testSuite) createProject() {
// 	b := JSON{
// 		"name":     s.testData.projectName,
// 		"password": s.testData.projectPW,
// 	}
// 	w := s.requestWithBody("POST", "/api/projects", b)

// 	response := &db.Project{}
// 	s.assertEqual(http.StatusCreated, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), response)
// 	s.assertTrue(response.ID > 0)
// 	s.assertEqual(s.testData.projectName, response.Name)
// }

// func (s *testSuite) listProjects() {
// 	w := s.request("GET", "/api/projects")

// 	var response []db.Project
// 	s.assertEqual(http.StatusOK, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), &response)
// 	s.assertTrue(len(response) == 1)
// 	s.assertEqual(response[0].Name, s.testData.projectName)
// 	s.testData.projectID = response[0].ID
// 	s.testData.projectIDString = strconv.Itoa(int(response[0].ID))
// }

// func (s *testSuite) selectProject() {
// 	b := JSON{
// 		"projectId": s.testData.projectID,
// 		"password":  s.testData.projectPW,
// 	}
// 	w := s.requestWithBody("POST", "/api/login", b)

// 	type token struct {
// 		Code   int       `json:"code"`
// 		Expire time.Time `json:"expire"`
// 		Token  string    `json:"token"`
// 	}
// 	response := &token{}
// 	s.assertEqual(http.StatusOK, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), response)
// 	s.assertNotNil(response.Expire)
// 	s.assertNotEmpty(response.Token)
// 	s.token = response.Token
// }

// func (s *testSuite) createPlayers() {
// 	for _, p := range s.testData.players {
// 		b := JSON{
// 			"name": p.name,
// 		}
// 		w := s.requestWithBody("POST", "/api/projects/"+s.testData.projectIDString+"/players", b)
// 		s.assertEqual(http.StatusCreated, w.Code)
// 	}
// }

// func (s *testSuite) listPlayers() {
// 	w := s.request("GET", "/api/projects/1/players")

// 	var response []db.Player
// 	s.assertEqual(http.StatusOK, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), &response)
// 	s.assertTrue(len(response) == len(s.testData.players))
// 	for i, r := range response {
// 		s.assertEqual(r.Name, s.testData.players[i].name)
// 	}
// }

// func (s *testSuite) createGame() {
// 	b := JSON{
// 		"name": s.testData.gameName,
// 	}
// 	w := s.requestWithBody("POST", "/api/projects/"+s.testData.projectIDString+"/games", b)
// 	s.assertEqual(http.StatusCreated, w.Code)
// }

// func (s *testSuite) listGames() {
// 	w := s.request("GET", "/api/projects/"+s.testData.projectIDString+"/games")

// 	var response []db.Game
// 	s.assertEqual(http.StatusOK, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), &response)
// 	s.assertTrue(len(response) == 1)
// 	s.assertEqual(response[0].Name, s.testData.gameName)
// 	s.testData.gameID = response[0].ID
// }

// func (s *testSuite) createMatches() {
// 	for _, m := range s.testData.matches {
// 		var teams []JSON
// 		for _, t := range m.teams {
// 			teams = append(teams, JSON{
// 				"playerIds": t.players,
// 				"score":     t.score,
// 			})
// 		}
// 		body := JSON{"teams": teams}

// 		gameID := strconv.Itoa(int(s.testData.gameID))
// 		w := s.requestWithBody("POST", "/api/projects/"+s.testData.projectIDString+"/games/"+gameID+"/matches", body)

// 		response := &db.Match{}
// 		s.assertEqual(http.StatusCreated, w.Code)
// 		s.mustUnmarshal(w.Body.Bytes(), response)
// 		s.assertTrue(response.ID > 0)
// 		s.assertEqual(response.GameID, s.testData.gameID)
// 	}
// }

// func (s *testSuite) listMatches() {
// 	gameID := strconv.Itoa(int(s.testData.gameID))
// 	w := s.request("GET", "/api/projects/"+s.testData.projectIDString+"/games/"+gameID+"/matches")

// 	var response []compelo.MatchData
// 	s.assertEqual(http.StatusOK, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), &response)
// 	s.assertTrue(len(response) == len(s.testData.matches))

// 	// Sort results by ID so that they are
// 	// in the same order as the test data.
// 	sort.Slice(response, func(i, j int) bool {
// 		return response[i].ID < (response[j].ID)
// 	})

// 	for i, expectedMatch := range s.testData.matches {
// 		actualMatch := response[i]
// 		s.assertEqual(len(expectedMatch.teams), len(actualMatch.Teams))

// 		for j, expectedTeam := range expectedMatch.teams {
// 			actualTeam := actualMatch.Teams[j]
// 			s.assertEqual(len(expectedTeam.players), len(actualTeam.Players))
// 			s.assertEqual(expectedTeam.score, actualTeam.Score)
// 			s.assertEqual(expectedTeam.result, actualTeam.Result)
// 		}
// 	}
// }

// func (s *testSuite) loadPlayerStats() {
// 	gameID := strconv.Itoa(int(s.testData.gameID))
// 	w := s.request("GET", "/api/projects/"+s.testData.projectIDString+"/games/"+gameID+"/player-stats")

// 	var response []compelo.PlayerStats
// 	s.assertEqual(http.StatusOK, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), &response)
// 	s.assertTrue(len(response) == len(s.testData.players))

// 	for _, v := range response {
// 		s.assertEqual(s.testData.stats[v.Name].rating, v.Current.Rating)
// 		s.assertEqual(s.testData.stats[v.Name].peakRating, v.Current.PeakRating)
// 		s.assertEqual(s.testData.stats[v.Name].lowestRating, v.Current.LowestRating)
// 		s.assertEqual(s.testData.stats[v.Name].gameCount, v.Current.GameCount)
// 		s.assertEqual(s.testData.stats[v.Name].winCount, v.Current.WinCount)
// 		s.assertEqual(s.testData.stats[v.Name].drawCount, v.Current.DrawCount)
// 		s.assertEqual(s.testData.stats[v.Name].lossCount, v.Current.LossCount)
// 	}
// }

// func (s *testSuite) loadGameStats() {
// 	gameID := strconv.Itoa(int(s.testData.gameID))
// 	w := s.request("GET", "/api/projects/"+s.testData.projectIDString+"/games/"+gameID+"/game-stats")

// 	var response compelo.GameStats
// 	s.assertEqual(http.StatusOK, w.Code)
// 	s.mustUnmarshal(w.Body.Bytes(), &response)

// 	// Only three test games so far, all should appear in stats.
// 	s.assertTrue(len(response.MaxScoreDiff) == len(s.testData.matches))
// 	s.assertTrue(len(response.MaxScoreSum) == len(s.testData.matches))

// 	// The third test-game, 1v1v1 (4:4:2)
// 	s.assertTrue(response.MaxScoreSum[0].ID == 3)

// 	// The second test-game, 1v1 (4:1)
// 	s.assertTrue(response.MaxScoreDiff[0].ID == 2)
// }

// // ------ Helpers ------

// func (s *testSuite) requestWithBody(method, path string, body JSON) *httptest.ResponseRecorder {
// 	b, err := json.Marshal(body)
// 	if err != nil {
// 		s.testing.Error(err)
// 	}
// 	req, err := http.NewRequest(method, path, bytes.NewBuffer(b))
// 	if err != nil {
// 		s.testing.Error(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	if s.token != "" {
// 		req.Header.Set("Authorization", "Bearer "+s.token)
// 	}
// 	w := httptest.NewRecorder()
// 	s.handler.ServeHTTP(w, req)
// 	return w
// }

// func (s *testSuite) request(method, path string) *httptest.ResponseRecorder {
// 	req, err := http.NewRequest(method, path, nil)
// 	if err != nil {
// 		s.testing.Error(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	if s.token != "" {
// 		req.Header.Set("Authorization", "Bearer "+s.token)
// 	}
// 	w := httptest.NewRecorder()
// 	s.handler.ServeHTTP(w, req)
// 	return w
// }

// func (s *testSuite) mustUnmarshal(bytes []byte, target interface{}) {
// 	err := json.Unmarshal(bytes, target)
// 	if err != nil {
// 		s.testing.Error(err)
// 	}
// }

// func (s *testSuite) assertEqual(expected, actual interface{}) {
// 	assert.Equal(s.testing, expected, actual)
// }

// func (s *testSuite) assertTrue(value bool) {
// 	assert.True(s.testing, value)
// }

// func (s *testSuite) assertNotNil(obj interface{}) {
// 	assert.NotNil(s.testing, obj)
// }

// func (s *testSuite) assertNotEmpty(obj interface{}) {
// 	assert.NotEmpty(s.testing, obj)
// }
