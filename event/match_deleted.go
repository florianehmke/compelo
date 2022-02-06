package event

const MatchDeletedType EventType = "MatchDeleted"

// MatchDeleted event.
type MatchDeleted struct {
	EventMetaData
	GUID        string `json:"guid"`
	GameGUID    string `json:"gameGuid"`
	ProjectGUID string `json:"projectGuid"`
}

func (e *MatchDeleted) EventType() EventType {
	return MatchDeletedType
}
