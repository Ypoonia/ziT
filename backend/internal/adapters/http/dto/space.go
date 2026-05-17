package dto

import "time"

type CreateSpaceRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ParentID    *string `json:"parent_id,omitempty"`
}

type UpdateSpaceRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ParentID    *string `json:"parent_id,omitempty"`
}

type SpaceResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ParentID    *string    `json:"parent_id,omitempty"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	CreatedOn   *time.Time `json:"created_on,omitempty"`
	UpdatedBy   *string    `json:"updated_by,omitempty"`
	UpdatedOn   *time.Time `json:"updated_on,omitempty"`
}
