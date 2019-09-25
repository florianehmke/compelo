package match

import (
	"sort"
	"time"

	"compelo/db"
	"compelo/game"
	"compelo/player"
	"compelo/rating"
)

type Service struct {
	repository    Repository
	playerService *player.Service
	gameService   *game.Service
}

func NewService(db *db.DB, ps *player.Service, gs *game.Service) *Service {
	return &Service{
		repository:    repository{db},
		playerService: ps,
		gameService:   gs,
	}
}

func (s *Service) createMatch(param createMatchParameter, game game.Game) (Match, error) {
	param.gameID = game.ID
	param.date = time.Now()
	param.determineWinner()
	param.calculateTeamElo(s.playerService, game.ID)

	// Create the match.
	m, err := s.repository.create(param)
	if err != nil {
		return Match{}, err
	}

	// Update the ratings of the players who participated.
	if err := s.updatePlayerRatings(param, game); err != nil {
		return Match{}, err
	}

	return m, nil
}

func (s *Service) updatePlayerRatings(param createMatchParameter, game game.Game) error {
	for _, t := range param.Teams {
		for _, playerId := range t.PlayerIDs {
			_, err := s.playerService.UpdateRating(uint(playerId), game.ID, t.ratingDelta)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *createMatchParameter) calculateTeamElo(playerService *player.Service, gameID uint) {
	rm := rating.NewRatedMatch()
	for i, t := range p.Teams {
		sum := 0
		for _, pid := range t.PlayerIDs {
			r, err := playerService.LoadRating(uint(pid), gameID)
			if err != nil {
				return
			}
			sum += r.Rating
		}
		avg := sum / len(t.PlayerIDs)

		// The rating service expects a "rank" to sort players.
		// Here we just use the negative score instead, should
		// result in the same thing for most games..
		rm.AddPlayer(i, -t.Score, avg)
	}
	rm.Calculate()

	for i := range p.Teams {
		p.Teams[i].ratingDelta = rm.GetRatingDelta(i)
	}
}

func (p *createMatchParameter) determineWinner() {
	highScore := 0
	highScoreCount := 0
	for _, t := range p.Teams {
		if t.Score > highScore {
			highScore = t.Score
			highScoreCount = 1
		} else if t.Score == highScore {
			highScoreCount += 1
		}
	}
	if highScoreCount < len(p.Teams) {
		for i := range p.Teams {
			if p.Teams[i].Score == highScore {
				p.Teams[i].result = Win.String()
			} else {
				p.Teams[i].result = Loss.String()
			}
		}
	} else {
		for i := range p.Teams {
			p.Teams[i].result = Draw.String()
		}
	}
}

type MatchData struct {
	ID     uint       `json:"id"`
	Date   time.Time  `json:"date"`
	GameID uint       `json:"gameId"`
	Teams  []TeamData `json:"teams"`
}

type TeamData struct {
	ID          uint         `json:"id"`
	MatchID     uint         `json:"matchId"`
	Score       int          `json:"score"`
	Result      string       `json:"result"`
	RatingDelta int          `json:"ratingDelta"`
	Players     []PlayerData `json:"players"`
}

type PlayerData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

func (s *Service) LoadMatchesByGameID(gameID uint) ([]MatchData, error) {
	var matchDataList []MatchData

	matches, err := s.repository.loadByGameID(gameID)
	for _, match := range matches {
		matchData, err := s.LoadMatchByID(match.ID)
		if err != nil {
			return matchDataList, err
		}
		matchDataList = append(matchDataList, matchData)
	}

	// Sort matches by date.
	sort.Slice(matchDataList, func(i, j int) bool {
		return matchDataList[i].Date.After(matchDataList[j].Date)
	})

	return matchDataList, err
}

func (s *Service) LoadMatchByID(id uint) (MatchData, error) {
	// 1. Get basic match data.
	match, err := s.repository.loadByID(id)
	if err != nil {
		return MatchData{}, err
	}
	matchData := MatchData{
		ID:     match.ID,
		Date:   match.Date,
		GameID: match.GameID,
	}

	// 2. Get data about teams.
	teams, err := s.repository.loadTeamsByMatchID(id)
	if err != nil {
		return MatchData{}, err
	}
	for _, t := range teams {
		teamData := TeamData{
			ID:          t.ID,
			MatchID:     t.MatchID,
			Score:       t.Score,
			Result:      t.Result,
			RatingDelta: t.RatingDelta,
		}

		// 3. Get data about players.
		players, err := s.loadPlayersByMatchIDAndTeamID(id, t.ID)
		if err != nil {
			return MatchData{}, err
		}
		for _, p := range players {
			playerData := PlayerData{
				ID:        p.ID,
				Name:      p.Name,
				ProjectID: p.ProjectID,
			}
			teamData.Players = append(teamData.Players, playerData)
		}
		matchData.Teams = append(matchData.Teams, teamData)
	}

	// 4. Sort teams by score.
	sort.Slice(matchData.Teams, func(i, j int) bool {
		if matchData.Teams[i].Score == matchData.Teams[j].Score {
			return matchData.Teams[i].ID < matchData.Teams[j].ID
		}
		return matchData.Teams[i].Score > matchData.Teams[j].Score
	})

	return matchData, err
}

func (s *Service) loadPlayersByMatchIDAndTeamID(matchID, teamID uint) ([]player.Player, error) {
	appearances, err := s.repository.loadAppearancesByMatchIDAndTeamID(matchID, teamID)
	if err != nil {
		return nil, err
	}

	var players []player.Player
	for _, appearance := range appearances {
		p, err := s.playerService.LoadPlayerByID(appearance.PlayerID)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, err
}
