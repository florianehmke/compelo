package query

type Project struct {
	MetaData
	GUID string `json:"guid"`

	Name         string `json:"name"`
	PasswordHash []byte `json:"passwordHash"`

	players map[string]*Player
	games   map[string]*Game
}
