package event_test

import (
	"events_example/domain/event"
	"events_example/graph/model"
	"testing"

	"github.com/google/uuid"
)

var tr *event.TestRepository
var events map[uuid.UUID]model.Event

func TestGetEvent(t *testing.T) {
	newUUID := uuid.New()
	testEvent := model.Event{
		ID: newUUID, Name: "foo",
	}
	events = map[uuid.UUID]model.Event{
		newUUID: testEvent,
	}
	tr := event.NewTestRepository(&events)
	returnedEvent, err := tr.GetEvent(newUUID)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *returnedEvent != testEvent {
		t.Errorf("the wrong event was returned")
	}
}
