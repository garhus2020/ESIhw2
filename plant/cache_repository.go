package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

type CacheRepository struct {
	db *redis.Client
}

func NewCacheRepository(db *redis.Client) *CacheRepository {
	return &CacheRepository{
		db: db,
	}
}

func (r *CacheRepository) Create(request *Request) (*Request, error) {
	requestID := uuid.NewString()
	request.ID = requestID
	request.CreatedAt = time.Now()
	_, err := r.db.HSetNX(context.Background(), "app:request", requestID, request).Result()
	if err != nil {
		return nil, fmt.Errorf("create: redis error: %w", err)
	}

	return request, nil
}

func (r *CacheRepository) GetAll() ([]*Request, error) {

	res, err := r.db.HGetAll(context.Background(), "app:request").Result()
	if err != nil {
		return nil, fmt.Errorf("error getting requests, err: %v", err)
	}
	requests := []*Request{}
	for _, stringValue := range res {
		b := &Request{}
		err := json.Unmarshal([]byte(stringValue), b)
		if err != nil {
			return nil, fmt.Errorf("error decoding result, err: %v", err)
		}
		requests = append(requests, b)
	}

	return requests, nil
}
