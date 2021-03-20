package domain

import "time"

type Request struct {
	ID        interface{}
	Body      string
	CreatedAt time.Time
}

