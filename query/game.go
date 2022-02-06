package query

import "sort"

type Game struct {
	MetaData
	GUID        string `json:"guid"`
	ProjectGUID string `json:"projectGuid"`

	Name string `json:"name"`

	eloMatchList *eloMatchList
	playerStats  map[string]*PlayerStats
	gameStats    *GameStats
	competitions map[string]*Competition
}

func sortGamesByCreatedDate(values []*Game) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].ID < values[j].ID
	})
}
