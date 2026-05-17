package domain

import "time"

type TicketStatus string

const (
	TicketStatusTodo       TicketStatus = "todo"
	TicketStatusInProgress TicketStatus = "in_progress"
	TicketStatusDone       TicketStatus = "done"
)

type TicketPriority string

const (
	TicketPriorityLow    TicketPriority = "low"
	TicketPriorityMedium TicketPriority = "medium"
	TicketPriorityHigh   TicketPriority = "high"
)

type Ticket struct {
	ID          string
	Title       string
	Description string
	Status      TicketStatus
	ParentID    string
	Priority    *TicketPriority
	Assignee    *string
	Reporter    *string
	CreatedOn   *time.Time
	UpdatedBy   *string
	UpdatedOn   *time.Time
}
