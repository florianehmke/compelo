package query

import (
	"compelo/event"
	"log"
	"sync"
	"time"
)

type Service struct {
	sync.RWMutex
	bus *event.Bus

	data *data

	defaultHandler     *defaultHandler
	gameStatsHandler   *gameStatsHandler
	playerStatsHandler *playerStatsHandler
}

func NewService(bus *event.Bus) *Service {
	data := &data{projects: make(map[string]*Project)}

	svc := Service{
		bus:                bus,
		data:               data,
		defaultHandler:     &defaultHandler{data: data},
		gameStatsHandler:   &gameStatsHandler{data: data},
		playerStatsHandler: &playerStatsHandler{data: data},
	}

	channel := bus.Subscribe()
	go func() {
		for event := range channel {
			svc.on(event)
		}
	}()

	return &svc
}

func (svc *Service) on(e event.Event) {
	defer svc.bus.MessageProcessed()

	svc.Lock()
	defer svc.Unlock()

	log.Println("[query] handling event", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.ProjectCreated:
		svc.defaultHandler.handleProjectCreated(e)
	case *event.GameCreated:
		svc.defaultHandler.handleGameCreated(e)
	case *event.PlayerCreated:
		svc.defaultHandler.handlePlayerCreated(e)
	case *event.MatchCreated:
		svc.defaultHandler.handleMatchCreated(e)
		svc.gameStatsHandler.handleMatchCreated(e)
		svc.playerStatsHandler.handleMatchCreated(e)
	}
}

// MetaData contains common meta data for query objects.
type MetaData struct {
	ID          uint64    `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

func (md *MetaData) getCreatedDate() time.Time {
	return md.CreatedDate
}

func (md *MetaData) getUpdatedDate() time.Time {
	return md.getUpdatedDate()
}

type defaultHandler struct {
	data *data
}

type data struct {
	projects map[string]*Project
}
