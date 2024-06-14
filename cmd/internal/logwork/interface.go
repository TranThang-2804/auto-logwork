package logwork

import (
	"time"
)

type ProjectTracking interface {
	GetTicketToLog() ([]Ticket, error)
	GetDayToLog() ([]time.Time, error)
	LogWork(ticket []Ticket, day []time.Time) error
}

type Ticket struct {
	ID string
}
