package logwork

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

func (j *Jira) GetTicketToLog() ([]Ticket, error) {
  return []Ticket{}, nil
}

func (j *Jira) GetDayToLog() (string, error) {
  return "", nil
}

func (j *Jira) LogWork(ticket []Ticket, day string) error {
  return nil
}
