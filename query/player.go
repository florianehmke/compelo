package query

type Player struct {
	MetaData
	GUID        string `json:"guid"`
	ProjectGUID string `json:"projectGuid"`

	Name string `json:"name"`

	ratings map[string]*Rating
}
