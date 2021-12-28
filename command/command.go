package command

import (
	"compelo/event"
	"fmt"
	"log"
	"sync"
	"time"
)

type Service struct {
	uniqueConstraints
	data *data

	sync.RWMutex
	changes []event.Event
	version int
	store   *event.Store
}

type Response struct {
	GUID string `json:"guid"`
}

func NewService(store *event.Store, events []event.Event) *Service {
	p := &Service{
		data: &data{
			projects: make(map[string]project),
		},
		store: store,
	}

	for _, event := range events {
		p.on(event)
	}

	return p
}

func (svc *Service) on(e event.Event) {
	log.Println("[command] handling event", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.ProjectCreated:
		svc.handleProjectCreated(e)
	case *event.GameCreated:
		svc.handleGameCreated(e)
	case *event.PlayerCreated:
		svc.handlePlayerCreated(e)
	case *event.MatchCreated:
		svc.handleMatchCreated(e)
	}
	svc.version++
}

func (svc *Service) raise(event event.Event) error {
	event.SetDate(time.Now())

	svc.changes = append(svc.changes, event)
	svc.on(event)

	if err := svc.store.StoreEvent(event); err != nil {
		return fmt.Errorf("storing event failed: %w", err)
	}

	return nil
}

type data struct {
	projects map[string]project
}
