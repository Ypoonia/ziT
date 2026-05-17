package dto

import "time"

type CreateIterationRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    string    `json:"parent_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type UpdateIterationRequest struct {
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
}

type IterationResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    string     `json:"parent_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	CreatedOn   *time.Time `json:"created_on,omitempty"`
	UpdatedBy   *string    `json:"updated_by,omitempty"`
	UpdatedOn   *time.Time `json:"updated_on,omitempty"`
}
