package graph

import "events_example/domain/event"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TR *event.TestRepository
}
