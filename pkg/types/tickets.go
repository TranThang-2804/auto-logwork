package types

import "time"

type Ticket struct {
	ID        string
	Summary   string
	TimeToLog int64
	DateToLog time.Time
}
