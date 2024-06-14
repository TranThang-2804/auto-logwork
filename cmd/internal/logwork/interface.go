package logwork

import (
	"github.com/TranThang-2804/auto-logwork/pkg/types"
  "time"
)

type ProjectTracking interface {
	GetTicketToLog(config *types.Config) ([]Ticket, error)
	GetDayToLog(config *types.Config) ([]time.Time, error)
	LogWork(config *types.Config, ticket []Ticket, day []time.Time) error
}

type Ticket struct {
	ID string
}
