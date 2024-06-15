package logwork

import (
	"slices"
	"time"

	"github.com/TranThang-2804/auto-logwork/pkg/types"
)

func defaultLogWorkAlgorithm(ticket []types.Ticket, logworkList []types.LogWorkStatus) ([]types.LogAction, error) {
	defaultShiftTime := 8
	workingDay := []int{1, 2, 3, 4, 5}
  startShiftHour := 8*time.Hour + 30*time.Minute  // 8h30

	logActionList := []types.LogAction{}

	for i := range logworkList {
		if logworkList[i].TimeSpent/3600 < int64(defaultShiftTime) && slices.Contains(workingDay, i) {
			ticketIndex := i % len(ticket)
			logActionList = append(logActionList, types.LogAction{
				TimeToLog:   int64(defaultShiftTime*3600) - logworkList[i].TimeSpent,
				TicketToLog: ticket[ticketIndex],
				DateToLog:   logworkList[i].Date.Add(time.Duration(startShiftHour)),
			})
		}
	}

	return logActionList, nil
}
