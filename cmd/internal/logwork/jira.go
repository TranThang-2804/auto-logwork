package logwork

import (
	"fmt"
	"log"
	"time"

	"github.com/andygrunwald/go-jira"
)

type Jira struct {
	endpoint string
	userName string
	apiToken string
}

func NewJira(endpoint string, userName string, apiToken string) *Jira {
	return &Jira{
		endpoint:   endpoint,
    userName:   userName,
    apiToken:   apiToken,
	}
}

func (j *Jira) GetTicketToLog() ([]Ticket, error) {
	tp := jira.BasicAuthTransport{
		Username: j.userName,
		Password: j.apiToken,
	}

	client, err := jira.NewClient(tp.Client(), j.endpoint)
	if err != nil {
		log.Fatalf("Error creating JIRA client: %v", err)
	}

	// JQL query to fetch your tickets. Customize this query as needed.
	jql := fmt.Sprintf(`assignee = "%s" AND status != Closed ORDER BY created DESC`, j.userName)

	issues, _, err := client.Issue.Search(jql, &jira.SearchOptions{
		MaxResults: 10, // Adjust the number of results as needed
	})
	if err != nil {
		log.Fatalf("Error fetching JIRA issues: %v", err)
	}

	// Print the fetched issues
	for _, issue := range issues {
		fmt.Printf("Issue: %s, Summary: %s, Status: %s\n", issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
	}
	return []Ticket{}, nil
}

func (j *Jira) GetDayToLog() ([]time.Time, error) {
	return []time.Time{}, nil
}

func (j *Jira) LogWork(ticket []Ticket, day []time.Time) error {
	fmt.Println("Log work")
	return nil
}
