package logwork

import (
	"fmt"
	"log"
	"time"

	"github.com/TranThang-2804/auto-logwork/pkg/types"
	"github.com/andygrunwald/go-jira"
)

type Jira struct {
	endpoint string
	userName string
	apiToken string
	client   *jira.Client
}

func NewJira(endpoint string, userName string, apiToken string) *Jira {
	tp := jira.BasicAuthTransport{
		Username: userName,
		Password: apiToken,
	}

	client, err := jira.NewClient(tp.Client(), endpoint)

	if err != nil {
		log.Fatalf("Error creating JIRA client: %v", err)
	}

	return &Jira{
		endpoint: endpoint,
		userName: userName,
		apiToken: apiToken,
		client:   client,
	}
}

func (j *Jira) GetTicketToLog() ([]types.Ticket, error) {
	// JQL query to fetch your tickets. Customize this query as needed.
	jql := fmt.Sprintf(`assignee = "%s" AND status IN (Resolved, "In Progress") AND type != Epic ORDER BY created DESC`, j.userName)

	ticketList := []types.Ticket{}

	issues, _, err := j.client.Issue.Search(jql, &jira.SearchOptions{
		MaxResults: 10, // Adjust the number of results as needed
	})
	if err != nil {
		log.Fatalf("Error fetching JIRA issues: %v", err)
	}

	// Print the fetched issues
	for _, issue := range issues {
		fmt.Printf("Issue: %s, Summary: %s, Status: %s\n", issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
		ticketList = append(ticketList, types.Ticket{
			ID:      issue.Key,
			Summary: issue.Fields.Summary,
		})
	}
	return ticketList, nil
}

func (j *Jira) GetDayToLog() ([]types.LogWorkStatus, error) {
	// Calculate the start of the current week (Monday)
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	var startOfWeek time.Time

	// Sunday is 0 -> we need to handle this
	if now.Weekday().String() == "Sunday" {
		startOfWeek = start.AddDate(0, 0, -6) // Adjust according to your week's start day
	} else {
		startOfWeek = start.AddDate(0, 0, -int(now.Weekday()+1)) // Adjust according to your week's start day
	}

	fmt.Println("Start of week: ", startOfWeek)

	logworkList := make([]types.LogWorkStatus, 7)

	// Create the correct date
	for i := range logworkList {
		if i == 0 {
			logworkList[i].Date = startOfWeek.AddDate(0, 0, 6)
		} else {
			logworkList[i].Date = startOfWeek.AddDate(0, 0, i-1)
		}
	}

	// JQL query to fetch issues assigned to you
	jql := fmt.Sprintf(`assignee = "%s" ORDER BY created DESC`, j.userName)

	issues, _, err := j.client.Issue.Search(jql, &jira.SearchOptions{
		MaxResults: 100, // Adjust the number of results as needed
	})
	if err != nil {
		log.Fatalf("Error fetching JIRA issues: %v", err)
	}

	fmt.Println("Work logs for the current week:")

	for _, issue := range issues {
		worklogs, _, err := j.client.Issue.GetWorklogs(issue.Key)
		if err != nil {
			log.Printf("Error fetching worklogs for issue %s: %v", issue.Key, err)
			continue
		}

		for _, worklog := range worklogs.Worklogs {
			worklogTimeStarted, _ := worklog.Started.MarshalJSON()
			worklogTime, err := time.Parse("\"2006-01-02T15:04:05.999-0700\"", string(worklogTimeStarted))
			if err != nil {
				log.Printf("Error parsing worklog time for issue %s: %v", issue.Key, err)
				continue
			}

			if worklogTime.After(startOfWeek) {
				logworkList[worklogTime.Weekday()].Add(int64(worklog.TimeSpentSeconds))
			}
		}
	}

	for i := range logworkList {
		fmt.Printf("Day: %d, Time Spent: %d Hours\n", i, logworkList[i].TimeSpent/3600)
	}

	return logworkList, nil
}

func (j *Jira) LogWork(ticket []types.Ticket, logworkList []types.LogWorkStatus) error {
	logActionList, _ := defaultLogWorkAlgorithm(ticket, logworkList)

  fmt.Println("----------------Ticket to log-------------------")
	for i := range logActionList {
    fmt.Printf("Ticket ID: %s, Tiket Summary: %s, Time to log: %s, Date to log: %s\n", logActionList[i].TicketToLog.ID, logActionList[i].TicketToLog.Summary, logActionList[i].DateToLog)
	}

    

	return nil
}
