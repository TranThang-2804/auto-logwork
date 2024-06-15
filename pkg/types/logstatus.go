package types

import (
	"log"
	"time"
  "errors"
)

type LogWorkStatus struct {
	Date      time.Time
	TimeSpent time.Duration
}

func (l *LogWorkStatus) New(date time.Time, timeSpent time.Duration) (*LogWorkStatus, error ) {
  return &LogWorkStatus{
    Date: date,
    TimeSpent: timeSpent,
  }, nil
}

func (l *LogWorkStatus) Add(timeSpent time.Duration) error {
  if (l == nil) {
    log.Fatal("LogWorkStatus is nil")
    return errors.New("LogWorkStatus is nil")
  }

  l.TimeSpent = time.Duration(l.TimeSpent.Seconds()+ timeSpent.Seconds())

  return nil
} 
