package logwork

import (
	"time"
  "github.com/TranThang-2804/auto-logwork/pkg/types"
)

type ProjectTracking interface {
	GetTicketToLog() ([]types.Ticket, error)
	GetDayToLog() ([]time.Time, error)
	LogWork(ticket []types.Ticket, day []time.Time) error
}
