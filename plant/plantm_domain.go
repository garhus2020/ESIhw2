package main

import (
	"time"
)

type Plantm struct {
	ID        interface{} `bson:"_id,omitempty" json:"ID"`
	Ident     string
	Name      string
	Price     string
	Status    string
	CreatedAt time.Time
}
