package domain

import "time"

type Space struct {
	ID          string
	Name        string
	Description string
	ParentID    *string
	CreatedBy   *string
	CreatedOn   *time.Time
	UpdatedBy   *string
	UpdatedOn   *time.Time
}
