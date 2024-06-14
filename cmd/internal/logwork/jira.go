package logwork

import (
	"fmt"
	"log"
	"time"

	"github.com/andygrunwald/go-jira"
)

type Jira struct {
	endpoint   string
	credential string
}

func NewJira(endpoint string, credential string) *Jira {
	return &Jira{
		endpoint:   endpoint,
		credential: credential,
	}
}

func (j *Jira) GetTicketToLog() ([]Ticket, error) {
	// Replace with your JIRA domain, username, and API token
	jiraDomain := "https://your-jira-domain.atlassian.net"
	username := "your-email@example.com"
	apiToken := "your-api-token"

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: apiToken,
	}

	client, err := jira.NewClient(tp.Client(), jiraDomain)
	if err != nil {
		log.Fatalf("Error creating JIRA client: %v", err)
	}

	// JQL query to fetch your tickets. Customize this query as needed.
	jql := fmt.Sprintf("assignee = %s AND status != Closed ORDER BY created DESC", username)

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
