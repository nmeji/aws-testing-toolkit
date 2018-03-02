package event

import "time"

type Event struct {
	Type        string
	Attachment  interface{}
	TimeEmitted time.Time
}

func New(typeOfEvent string) *Event {
	return &Event{Type: typeOfEvent}
}
