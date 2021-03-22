package domain

import (
	"time"
)

type Plant struct {
	ID        int `json:"id"`
	Ident     string `json:"ident"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Status    string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
