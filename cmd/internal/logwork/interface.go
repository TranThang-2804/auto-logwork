package logwork

type ProjectTracking interface {
  GetTicketToLog() ([]Ticket, error)
  GetDayTolog() (string, error)
  LogWork(ticket Ticket) error
}

type Ticket struct {
  ID string
}
