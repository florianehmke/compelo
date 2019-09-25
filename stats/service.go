package stats

import (
	"sort"
	"time"

	"compelo/db"
	"compelo/player"
	"compelo/rating"
)

type Service struct {
	db *db.DB
	ps *player.Service
}

func NewService(db *db.DB, ps *player.Service) *Service {
	return &Service{db: db, ps: ps}
}

type Player struct {
	ID             uint     `json:"id"`
	Name           string   `json:"name"`
	ProjectID      uint     `json:"projectId"`
	Rating         int      `json:"rating"`
	PeakRating     int      `json:"peakRating"`
	LowestRating   int      `json:"lowestRating"`
	GameCount      int      `json:"gameCount"`
	WinCount       int      `json:"winCount"`
	DrawCount      int      `json:"drawCount"`
	LossCount      int      `json:"lossCount"`
	RatingProgress []Rating `json:"ratingProgress"`
}

type Rating struct {
	Rating int       `json:"rating"`
	Date   time.Time `json:"date"`
}

func (s *Service) LoadPlayerStatsByGameID(gameID uint) ([]Player, error) {
	var ratings []player.Rating
	err := s.db.Where(player.Rating{GameID: gameID}).Find(&ratings).Error
	if err != nil {
		return nil, err
	}

	var players []Player
	for _, r := range ratings {
		p, err := s.ps.LoadPlayerByID(r.PlayerID)
		if err != nil {
			return nil, err
		}

		pws := Player{
			ID:           p.ID,
			Name:         p.Name,
			ProjectID:    p.ProjectID,
			Rating:       r.Rating,
			PeakRating:   rating.InitialRating,
			LowestRating: rating.InitialRating,
		}
		if err := s.applyRatingStats(&pws); err != nil {
			return nil, err
		}
		if err := s.applyGameStats(&pws); err != nil {
			return nil, err
		}
		players = append(players, pws)
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Rating > (players[j].Rating)
	})
	return players, nil
}

func (s *Service) applyRatingStats(player *Player) (err error) {
	selectRatings := `
		SELECT m.date, t.rating_delta
		FROM matches m
				 JOIN appearances a ON m.id = a.match_id
				 JOIN teams t ON a.team_id = t.id
		WHERE a.player_id = ?
		ORDER BY m.date ASC`

	rows, err := s.db.Raw(selectRatings, player.ID).Rows()
	if err != nil {
		return err
	}
	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		var r Rating
		err := rows.Scan(&r.Date, &r.Rating)
		if err != nil {
			return err
		}
		player.RatingProgress = append(player.RatingProgress, r)
	}

	current := rating.InitialRating
	for i, v := range player.RatingProgress {
		current = current + v.Rating
		player.RatingProgress[i].Rating = current

		if current > player.PeakRating {
			player.PeakRating = current
		}
		if current < player.LowestRating {
			player.LowestRating = current
		}
	}
	return nil
}

func (s *Service) applyGameStats(player *Player) (err error) {
	selectRatings := `
		SELECT t.result
		FROM players p
				 JOIN appearances a ON p.id = a.player_id
				 JOIN teams t ON a.team_id = t.id
		WHERE p.id = ?`

	rows, err := s.db.Raw(selectRatings, player.ID).Rows()
	if err != nil {
		return err
	}
	defer func() {
		err = rows.Close()
	}()

	var results []string
	for rows.Next() {
		var r string
		err := rows.Scan(&r)
		if err != nil {
			return err
		}
		results = append(results, r)
	}

	wins := 0
	draws := 0
	losses := 0
	for _, r := range results {
		switch r {
		case "Win":
			wins += 1
		case "Draw":
			draws += 1
		case "Loss":
			losses += 1
		}
	}

	player.GameCount = len(results)
	player.WinCount = wins
	player.DrawCount = draws
	player.LossCount = losses
	return nil
}
