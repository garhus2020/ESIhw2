package main

import (
	"time"
)

type Plant struct {
	ID        int
	Ident     string
	Name      string
	Price     string
	Status    string
	CreatedAt time.Time
}
