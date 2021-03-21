package domain

import (
	"encoding/json"
	"time"
)

type Request struct {
	ID        interface{}
	Body      string
	CreatedAt time.Time
}

func (t *Request) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Request) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	return nil
}
