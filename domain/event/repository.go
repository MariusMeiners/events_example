package event

import (
	"errors"
	"events_example/graph/model"

	"github.com/google/uuid"
)

type TestRepository struct {
	db *map[uuid.UUID]model.Event
}

func NewTestRepository(db *map[uuid.UUID]model.Event) *TestRepository {
	return &TestRepository{
		db: db,
	}
}

func (tr *TestRepository) GetEvent(uuid uuid.UUID) (*model.Event, error) {
	var event model.Event
	elem, ok := (*tr.db)[uuid]
	if !ok {
		return &event, errors.New("element with given uuid does not exist")
	}
	event = elem
	return &event, nil
}

func (tr *TestRepository) GetEvents() []*model.Event {
	var events []*model.Event

	for _, v := range *tr.db {
		events = append(events, &v)
	}

	return events
}

func (tr *TestRepository) AddEvent(event model.Event) error {
	_, ok := (*tr.db)[event.ID]
	if !ok {
		return errors.New("event with given uuid already exists")
	}

	(*tr.db)[event.ID] = event
	return nil
}

func (tr *TestRepository) DeleteEvent(event model.Event) error {
	_, ok := (*tr.db)[event.ID]
	if !ok {
		return errors.New("element was not in map")
	}

	return nil
}
