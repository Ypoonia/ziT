package dto

import (
	"time"

	"github.com/Ypoonia/ziT/backend/internal/core/domain"
)

type CreateTicketRequest struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Status      domain.TicketStatus    `json:"status"`
	ParentID    string                 `json:"parent_id"`
	Priority    *domain.TicketPriority `json:"priority,omitempty"`
	Assignee    *string                `json:"assignee,omitempty"`
	Reporter    *string                `json:"reporter,omitempty"`
}

type UpdateTicketRequest struct {
	Title       *string                `json:"title,omitempty"`
	Description *string                `json:"description,omitempty"`
	Status      *domain.TicketStatus   `json:"status,omitempty"`
	Priority    *domain.TicketPriority `json:"priority,omitempty"`
	Assignee    *string                `json:"assignee,omitempty"`
}

type TicketResponse struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Status      domain.TicketStatus    `json:"status"`
	ParentID    string                 `json:"parent_id"`
	Priority    *domain.TicketPriority `json:"priority,omitempty"`
	Assignee    *string                `json:"assignee,omitempty"`
	Reporter    *string                `json:"reporter,omitempty"`
	CreatedOn   *time.Time             `json:"created_on,omitempty"`
	UpdatedBy   *string                `json:"updated_by,omitempty"`
	UpdatedOn   *time.Time             `json:"updated_on,omitempty"`
}
