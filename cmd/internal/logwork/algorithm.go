package logwork

import (
	"fmt"
	"slices"

	"github.com/TranThang-2804/auto-logwork/pkg/types"
)

func defaultLogWorkAlgorithm(ticket []types.Ticket, logworkList []types.LogWorkStatus) error {
	defaultShiftTime := 8
  workingDay := []int{1,2,3,4,5}

	for i := range logworkList {
		if logworkList[i].TimeSpent/3600 < int64(defaultShiftTime) && slices.Contains(workingDay, i) {
			ticketIndex := i % len(ticket)
			ticket[ticketIndex].TimeToLog = int64(defaultShiftTime*3600) - logworkList[i].TimeSpent
			ticket[ticketIndex].DateToLog = logworkList[i].Date
			fmt.Printf("Ticket ID: %s, Time to log: %d, Date to log: %s", ticket[i].ID, ticket[i].TimeToLog, ticket[i].DateToLog)
		}
	}

	return nil
}
