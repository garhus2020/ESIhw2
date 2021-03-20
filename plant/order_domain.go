package main

import (
	"time"
)

type Order struct {
	ID        int
	Ident     string
	Name      string
	Price     string
	Status    string
	Start     string
	End       string
	CreatedAt time.Time
}
