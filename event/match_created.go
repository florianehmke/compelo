package event

import (
	"time"
)

const MatchCreatedType EventType = "MatchCreated"

// MatchCreated event.
type MatchCreated struct {
	EventMetaData
	GUID             string    `json:"guid"`
	GameGUID         string    `json:"gameGuid"`
	ProjectGUID      string    `json:"projectGuid"`
	CompetitionGUID  string    `json:"competitionGUID"` // optional
	Date             time.Time `json:"date" ts_type:"string"`
	CompetitionRound int       `json:"competitionRound"` // optional
	CompetitionDay   int       `json:"competitionDay"`   // optional
	Teams            []struct {
		PlayerGUIDs []string
		Score       int
	} `json:"teams"`
}

func (e *MatchCreated) EventType() EventType {
	return MatchCreatedType
}
