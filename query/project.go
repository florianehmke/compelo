package query

import "sort"

type Project struct {
	MetaData
	GUID string `json:"guid"`

	Name         string `json:"name"`
	PasswordHash []byte `json:"passwordHash"`

	players map[string]*Player
	games   map[string]*Game
}

func sortProjectsByCreatedDate(values []*Project) {
	sort.Slice(values, func(i, j int) bool {
		return values[i].getCreatedDate().Before(values[j].getCreatedDate())
	})
}
