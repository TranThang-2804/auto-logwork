package logwork

type Logwork interface {
  GetAllTicket() ([]Ticket, error)
  GetDayTolog() (string, error)
  LogWork(ticket Ticket) error
}

type Ticket struct {
  ID string
}
