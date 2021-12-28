package query

import "time"

type PlayerStats struct {
	*Player                  // embedded player
	Current Stats            `json:"current"`
	History map[string]Stats `json:"history" ts_type:"{[key: string]: Stats}"`
}

type Stats struct {
	Rating       int `json:"rating"`
	PeakRating   int `json:"peakRating"`
	LowestRating int `json:"lowestRating"`
	GameCount    int `json:"gameCount"`
	WinCount     int `json:"winCount"`
	DrawCount    int `json:"drawCount"`
	LossCount    int `json:"lossCount"`
}

func (p *PlayerStats) addResult(match *Match, team *Team) {
	// rating (peak, low, current)
	p.Current.Rating = p.Current.Rating + team.RatingDelta
	if p.Current.Rating > p.Current.PeakRating {
		p.Current.PeakRating = p.Current.Rating
	}
	if p.Current.Rating < p.Current.LowestRating {
		p.Current.LowestRating = p.Current.Rating
	}

	// game results
	switch team.Result {
	case Win:
		p.Current.WinCount += 1
	case Draw:
		p.Current.DrawCount += 1
	case Loss:
		p.Current.LossCount += 1
	}
	p.Current.GameCount += 1
	p.History[formatDate(match.Date)] = p.Current
}

func (p *PlayerStats) copyCurrentResultToHistory(date time.Time) {
	p.History[formatDate(date)] = p.Current
}

func formatDate(date time.Time) string {
	withoutTime := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	return withoutTime.Format("2006-01-02")
}
