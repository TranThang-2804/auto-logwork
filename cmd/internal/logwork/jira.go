package logwork

import (
	"fmt"
	"time"

	"github.com/TranThang-2804/auto-logwork/pkg/types"
)

type Jira struct {
  endpoint string
  credential string
}

func NewJira(endpoint string, credential string) *Jira {
  return &Jira{
    endpoint: endpoint,
    credential: credential,
  }
}

func (j *Jira) GetTicketToLog(config *types.Config) ([]Ticket, error) {
  return []Ticket{}, nil
}

func (j *Jira) GetDayToLog(config *types.Config) ([]time.Time, error) {
  return []time.Time{}, nil
}

func (j *Jira) LogWork(config *types.Config, ticket []Ticket, day []time.Time) error {
  fmt.Println("Log work")
  return nil
}
