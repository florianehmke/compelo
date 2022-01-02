package query

import "sort"

type Competition struct {
	MetaData
	GUID        string `json:"guid"`
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`

	Name   string `json:"name"`
	Rounds int    `json:"rounds"`

	Teams []*CompetitionTeam `json:"teams"`
}

type CompetitionTeam struct {
	Players []*Player `json:"players"`
}

func sortCompetitionsByCreatedDate(values []*Competition) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].ID > values[j].ID
	})
}
