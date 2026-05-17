package dto

import "time"

type CreateUserRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       *int    `json:"age,omitempty"`
	Gender    *string `json:"gender,omitempty"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Age       *int    `json:"age,omitempty"`
	Gender    *string `json:"gender,omitempty"`
}

type UserResponse struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Age        *int       `json:"age,omitempty"`
	Gender     *string    `json:"gender,omitempty"`
	SignedUpOn *time.Time `json:"signed_up_on,omitempty"`
}
