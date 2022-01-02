package event

const CompetitionCreatedType EventType = "CompetitionCreated"

type CompetitionCreated struct {
	EventMetaData
	GUID        string `json:"guid"`
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`

	Name   string `json:"name"`
	Rounds int    `json:"rounds"`
	Teams  []struct {
		PlayerGUIDs []string
	} `json:"teams"`
}

func (e *CompetitionCreated) EventType() EventType {
	return CompetitionCreatedType
}
