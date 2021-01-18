package model

import "github.com/google/uuid"

type Event struct {
	ID   uuid.UUID
	Name string
	// Performers []Performer
}

// type Performer struct {
// 	ID   uuid.UUID
// 	Name string
// }
