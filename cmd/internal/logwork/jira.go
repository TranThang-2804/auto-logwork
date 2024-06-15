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
	jql := fmt.Sprintf(`assignee = "%s" AND status IN (Resolved, "In Progress", Closed) AND type != Epic ORDER BY created DESC`, j.userName)

	issues, _, err := j.client.Issue.Search(jql, &jira.SearchOptions{
		MaxResults: 10, // Adjust the number of results as needed
	})
	if err != nil {
		log.Fatalf("Error fetching JIRA issues: %v", err)
	}

	// Print the fetched issues
	for _, issue := range issues {
		fmt.Printf("Issue: %s, Summary: %s, Status: %s\n", issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
	}
	return []types.Ticket{}, nil
}

func (j *Jira) GetDayToLog() ([]time.Time, error) {
	// Calculate the start of the current week (Monday)
	now := time.Now()
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday())+1) // Adjust according to your week's start day
  fmt.Println("Start of week: ", startOfWeek)

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
				fmt.Printf("Issue: %s, Time Spent: %s, Started: %s\n", issue.Key, worklog.TimeSpent, string(worklogTimeStarted))
			}
		}
	}

	return []time.Time{}, nil
}

func (j *Jira) LogWork(ticket []types.Ticket, day []time.Time) error {
	fmt.Println("Log work")
	return nil
}
