package query

import "sort"

type Player struct {
	MetaData
	GUID        string `json:"guid"`
	ProjectGUID string `json:"projectGuid"`

	Name string `json:"name"`

	ratings map[string]*Rating
}

func sortPlayersByCreatedDate(values []*Player) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].ID < values[j].ID
	})
}
