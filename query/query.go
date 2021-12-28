package query

import (
	"compelo/event"
	"log"
	"sync"
	"time"
)

type Compelo struct {
	projects map[string]*Project

	sync.RWMutex
	bus *event.Bus
}

func New(bus *event.Bus) *Compelo {
	c := Compelo{
		projects: make(map[string]*Project),
		bus:      bus,
	}

	channel := bus.Subscribe()
	go func() {
		for event := range channel {
			c.on(event)
		}
	}()

	return &c
}

func (c *Compelo) on(e event.Event) {
	defer c.bus.MessageProcessed()

	c.Lock()
	defer c.Unlock()

	log.Println("Query handling event ", e.GetID(), e.EventType())

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
}

// MetaData contains common meta data for query objects.
type MetaData struct {
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

func (md *MetaData) getCreatedDate() time.Time {
	return md.CreatedDate
}

func (md *MetaData) getUpdatedDate() time.Time {
	return md.getUpdatedDate()
}
