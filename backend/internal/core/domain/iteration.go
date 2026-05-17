package domain

import "time"

type Iteration struct {
	ID          string
	Name        string
	Description string
	ParentID    string
	StartDate   time.Time
	EndDate     time.Time
	CreatedBy   *string
	CreatedOn   *time.Time
	UpdatedBy   *string
	UpdatedOn   *time.Time
}
