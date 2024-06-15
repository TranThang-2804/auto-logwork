package types

import "time"

type LogAction struct {
	TimeToLog     int64
	DateToLog     time.Time
	TicketToLog   Ticket
}
