package types

import (
	"errors"
	"log"
	"time"
)

type LogWorkStatus struct {
	Date      time.Time
	TimeSpent int64
}

func (l *LogWorkStatus) New(date time.Time, timeSpent int64) (*LogWorkStatus, error) {
	return &LogWorkStatus{
		Date:      date,
		TimeSpent: timeSpent,
	}, nil
}

func (l *LogWorkStatus) Add(timeSpent int64) error {
	if l == nil {
		log.Fatal("LogWorkStatus is nil")
		return errors.New("LogWorkStatus is nil")
	}

	l.TimeSpent = l.TimeSpent + timeSpent

	return nil
}
