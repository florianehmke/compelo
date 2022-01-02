package query

import "sort"

type Game struct {
	MetaData
	GUID        string `json:"guid"`
	ProjectGUID string `json:"projectGuid"`

	Name string `json:"name"`

	matches     map[string]*Match
	playerStats map[string]*PlayerStats
	gameStats   *GameStats
}

func sortGamesByCreatedDate(values []*Game) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].ID < values[j].ID
	})
}
