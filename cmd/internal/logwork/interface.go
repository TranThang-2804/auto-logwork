package logwork

type ProjectTracking interface {
  GetTicketToLog() ([]Ticket, error)
  GetDayToLog() (string, error)
  LogWork(ticket []Ticket, day string) error
}

type Ticket struct {
  ID string
}
