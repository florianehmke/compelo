package command

import (
	"compelo/event"
	"fmt"
	"log"
	"sync"
)

type Compelo struct {
	projects map[string]project
	uniqueConstraints

	sync.RWMutex
	changes []event.Event
	version int
	store   *event.Store
}

type Response struct {
	GUID string `json:"guid"`
}

func New(store *event.Store, events []event.Event) *Compelo {
	p := &Compelo{
		projects: make(map[string]project),
		store:    store,
	}

	for _, event := range events {
		p.on(event)
	}

	return p
}

func (c *Compelo) on(e event.Event) {
	log.Println("Command handling event ", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.ProjectCreated:
		c.handleProjectCreated(e)
	case *event.GameCreated:
		c.handleGameCreated(e)
	case *event.PlayerCreated:
		c.handlePlayerCreated(e)
	case *event.MatchCreated:
		c.handleMatchCreated(e)
	}
	c.version++
}

func (c *Compelo) raise(event event.Event) error {
	c.changes = append(c.changes, event)
	c.on(event)

	if err := c.store.StoreEvent(event); err != nil {
		return fmt.Errorf("storing event failed: %w", err)
	}

	return nil
}
