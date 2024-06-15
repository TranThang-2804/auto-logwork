package logwork

import (
  "github.com/TranThang-2804/auto-logwork/pkg/types"
)

type ProjectTracking interface {
	GetTicketToLog() ([]types.Ticket, error)
	GetDayToLog() ([]types.LogWorkStatus, error)
	LogWork(ticket []types.Ticket, logworkList []types.LogWorkStatus) error
}
