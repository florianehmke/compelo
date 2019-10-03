package db

import (
	"time"
)

type MatchResult struct {
	Date        time.Time
	RatingDelta int
	Result      Result
}

func (db *DB) LoadMatchResultsByPlayerIDAndGameID(playerID, gameID uint) (results []MatchResult, err error) {
	selectResults := `
		SELECT m.date, t.rating_delta, t.result
		FROM matches m
				 JOIN appearances a ON m.id = a.match_id
				 JOIN teams t ON a.team_id = t.id
		WHERE a.player_id = ? and m.game_id = ?
		ORDER BY m.date ASC
	`

	rows, err := db.gorm.Raw(selectResults, playerID, gameID).Rows()
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
	}()

	for rows.Next() {
		var result MatchResult
		err := rows.Scan(&result.Date, &result.RatingDelta, &result.Result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
