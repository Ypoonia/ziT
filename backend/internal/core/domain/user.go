package domain

import "time"

type User struct {
	ID         string
	FirstName  string
	LastName   string
	Age        *int
	Gender     *string
	SignedUpOn *time.Time
}
