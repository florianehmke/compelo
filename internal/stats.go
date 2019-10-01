package compelo

import (
	"sort"

	"compelo/pkg/rating"
)

type Player struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ProjectID    uint   `json:"projectId"`
	Rating       int    `json:"rating"`
	PeakRating   int    `json:"peakRating"`
	LowestRating int    `json:"lowestRating"`
	GameCount    int    `json:"gameCount"`
	WinCount     int    `json:"winCount"`
	DrawCount    int    `json:"drawCount"`
	LossCount    int    `json:"lossCount"`
}

func (svc *Service) LoadPlayerStatsByGameID(gameID uint) ([]Player, error) {
	ratings := svc.LoadRatingsByGameID(gameID)

	var players []Player
	for _, r := range ratings {
		p, err := svc.LoadPlayerByID(r.PlayerID)
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
		if err := svc.applyRatingStats(&pws); err != nil {
			return nil, err
		}
		if err := svc.applyGameStats(&pws); err != nil {
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
		SELECT t.rating_delta
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

	var deltas []int
	for rows.Next() {
		var delta int
		err := rows.Scan(&delta)
		if err != nil {
			return err
		}
		deltas = append(deltas, delta)
	}

	current := rating.InitialRating
	for _, delta := range deltas {
		current = current + delta
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
