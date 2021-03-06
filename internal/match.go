package compelo

import (
	"sort"
	"time"

	"github.com/pkg/errors"

	"compelo/internal/db"
	"compelo/pkg/rating"
)

type CreateMatchParameter struct {
	GameID uint
	Date   time.Time

	Teams []CreateMatchParameterTeam
}

type CreateMatchParameterTeam struct {
	PlayerIDs []int
	Score     int

	// Result + rating delta are determined by service.
	result      db.Result
	ratingDelta int
}

var (
	ErrTwoTeamsRequired      = errors.New("at least two teams required")
	ErrSameTeamSizeRequired  = errors.New("all teams need the same amount of players")
	ErrPlayerInMultipleTeams = errors.New("player can only be in one team")
)

func (p *CreateMatchParameter) validate() error {
	if len(p.Teams) < 2 {
		return ErrTwoTeamsRequired
	}
	teamSize := len(p.Teams[0].PlayerIDs)
	for _, t := range p.Teams {
		if len(t.PlayerIDs) != teamSize {
			return ErrSameTeamSizeRequired
		}
	}
	playerMap := map[int]bool{}
	for _, t := range p.Teams {
		for _, pid := range t.PlayerIDs {
			if _, ok := playerMap[pid]; ok {
				return ErrPlayerInMultipleTeams
			}
			playerMap[pid] = true
		}
	}
	return nil
}

func (svc *Service) CreateMatch(param CreateMatchParameter) (db.Match, error) {
	if err := param.validate(); err != nil {
		return db.Match{}, errors.Wrap(err, "create match parameter validation failed")
	}

	svc.determineResult(&param)
	if err := svc.calculateTeamElo(&param); err != nil {
		return db.Match{}, errors.Wrap(err, "could not calculate team elo")
	}

	var match db.Match
	if err := svc.db.DoInTransaction(func(tx db.Database) error {
		var err error

		// 1. Create match.
		if match, err = tx.CreateMatch(db.Match{GameID: param.GameID, Date: param.Date}); err != nil {
			return errors.Wrap(err, "create match failed")
		}

		// 2. Create teams.
		for _, team := range param.Teams {
			t, err := tx.CreateTeam(db.Team{
				MatchID:     match.ID,
				Score:       team.Score,
				Result:      team.result,
				RatingDelta: team.ratingDelta,
			})
			if err != nil {
				return errors.Wrap(err, "create team failed")
			}

			// 3. Create appearances for players.
			for _, playerID := range team.PlayerIDs {
				if _, err := tx.CreateAppearance(db.Appearance{
					MatchID:     match.ID,
					TeamID:      t.ID,
					PlayerID:    uint(playerID),
					RatingDelta: team.ratingDelta,
				}); err != nil {
					return errors.Wrap(err, "create appearance failed")
				}
			}
		}
		return nil
	}); err != nil {
		return db.Match{}, errors.Wrap(err, "create match in txn failed")
	}

	// Update the ratings of the players who participated.
	if err := svc.updatePlayerRatings(param); err != nil {
		return db.Match{}, errors.Wrap(err, "update player ratings failed")
	}

	return match, nil
}

func (svc *Service) updatePlayerRatings(param CreateMatchParameter) error {
	for _, t := range param.Teams {
		for _, playerID := range t.PlayerIDs {
			_, err := svc.UpdateRating(uint(playerID), param.GameID, t.ratingDelta)
			if err != nil {
				return errors.Wrapf(err, "update rating for player with id %d failed", playerID)
			}
		}
	}
	return nil
}

func (svc *Service) calculateTeamElo(param *CreateMatchParameter) error {
	rm := rating.NewRatedMatch()
	for i, t := range param.Teams {
		sum := 0
		for _, playerID := range t.PlayerIDs {
			r, err := svc.LoadOrCreateRatingByPlayerIDAndGameID(uint(playerID), param.GameID)
			if err != nil {
				return errors.Wrapf(err, "load rating for player with id %d failed", playerID)
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

	for i := range param.Teams {
		param.Teams[i].ratingDelta = rm.GetRatingDelta(i)
	}
	return nil
}

func (svc *Service) determineResult(param *CreateMatchParameter) {
	highScore := 0
	highScoreCount := 0
	for _, t := range param.Teams {
		if t.Score > highScore {
			highScore = t.Score
			highScoreCount = 1
		} else if t.Score == highScore {
			highScoreCount += 1
		}
	}
	if highScoreCount < len(param.Teams) {
		for i := range param.Teams {
			if param.Teams[i].Score == highScore {
				param.Teams[i].result = db.Win
			} else {
				param.Teams[i].result = db.Loss
			}
		}
	} else {
		for i := range param.Teams {
			param.Teams[i].result = db.Draw
		}
	}
}

type MatchData struct {
	ID     uint       `json:"id"`
	Date   time.Time  `json:"date" ts_type:"string"`
	GameID uint       `json:"gameId"`
	Teams  []TeamData `json:"teams"`
}

type TeamData struct {
	ID          uint         `json:"id"`
	MatchID     uint         `json:"matchId"`
	Score       int          `json:"score"`
	Result      db.Result    `json:"result"`
	RatingDelta int          `json:"ratingDelta"`
	Players     []PlayerData `json:"players"`
}

type PlayerData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ProjectID uint   `json:"projectId"`
}

func (svc *Service) LoadMatchesByGameID(gameID uint) ([]MatchData, error) {
	var matchDataList []MatchData

	matches := svc.db.LoadMatchesByGameID(gameID)
	for _, match := range matches {
		matchData, err := svc.LoadMatchByID(match.ID)
		if err != nil {
			return matchDataList, errors.Wrapf(err, "could not load match by id %d", match.ID)
		}
		matchDataList = append(matchDataList, matchData)
	}

	// Sort matches by date.
	sort.Slice(matchDataList, func(i, j int) bool {
		return matchDataList[i].Date.After(matchDataList[j].Date)
	})

	return matchDataList, nil
}

func (svc *Service) LoadMatchByID(id uint) (MatchData, error) {
	// 1. Get basic match data.
	match, err := svc.db.LoadMatchByID(id)
	if err != nil {
		return MatchData{}, err
	}
	matchData := MatchData{
		ID:     match.ID,
		Date:   match.Date,
		GameID: match.GameID,
	}

	// 2. Get data about teams.
	teams, err := svc.db.LoadTeamsByMatchID(id)
	if err != nil {
		return MatchData{}, errors.Wrapf(err, "could not load teams by match id %d", id)
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
		players, err := svc.db.LoadPlayersByMatchIDAndTeamID(id, t.ID)
		if err != nil {
			return MatchData{}, errors.Wrapf(err, "could not load players by match id %d", id)
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
