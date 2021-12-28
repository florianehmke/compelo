package query

type Game struct {
	MetaData
	GUID        string `json:"guid"`
	ProjectGUID string `json:"projectGuid"`

	Name string `json:"name"`

	matches map[string]*Match
}
