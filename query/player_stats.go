package query

type PlayerStats struct {
	Player                   // embedded player
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
