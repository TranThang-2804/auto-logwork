package logwork

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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
	fmt.Println("----------------Ticket able to log-------------------")
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
		startOfWeek = start.AddDate(0, 0, -int(now.Weekday())+1) // Adjust according to your week's start day
	}

	fmt.Println("----------------Your week worklog status-------------------")
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
    fmt.Printf("%s: Time Spent: %d Hours\n", time.Weekday(i), logworkList[i].TimeSpent/3600)
	}

	return logworkList, nil
}

func (j *Jira) LogWork(ticket []types.Ticket, logworkList []types.LogWorkStatus) error {
	logActionList, _ := defaultLogWorkAlgorithm(ticket, logworkList)

	fmt.Println("----------------Ticket to log-------------------")
	for i := range logActionList {
		fmt.Printf("Ticket ID: %s\tTiket Summary: %s\t\tTime to log: %dh\tDate to log: %s\n", logActionList[i].TicketToLog.ID, logActionList[i].TicketToLog.Summary, logActionList[i].TimeToLog/3600, logActionList[i].DateToLog)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("You sure to start logging work? [y/n]: ")
	status, _ := reader.ReadString('\n')
	status = status[:len(status)-1]

	if status == "n" {
		return nil
	} else if status != "y" {
		log.Println("Invalid input")
		return errors.New("Invalid input, valid input are y/n")
	}

	for i := range logActionList {
		worklog := &jira.WorklogRecord{
			Started:          (*jira.Time)(&logActionList[i].DateToLog),
			TimeSpentSeconds: int(logActionList[i].TimeToLog),
		}

		// Log work to the Jira issue
		_, response, err := j.client.Issue.AddWorklogRecord(logActionList[i].TicketToLog.ID, worklog)
		if err != nil {
			log.Fatalf("Failed to log work: %v", err)
		}
		defer response.Body.Close()

    fmt.Printf("Work logged to issue %s: %s successfully.\n", logActionList[i].TicketToLog.ID, logActionList[i].TicketToLog.Summary)
	}

	return nil
}
