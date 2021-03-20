package repository


import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/go-redis/redis/v8"
	"time"
	domain "github.com/garhus2020/ESIhw2/plant/pkg/domain"
)

type CacheRepository struct {
	db *redis.Client
}

func NewCacheRepository(db *redis.Client) *CacheRepository {
	return &CacheRepository{
		db: db,
	}
}

func (r *CacheRepository) Create(request *domain.Request) (*domain.Request, error) {
	requestID := uuid.NewString()
	request.ID = requestID
	request.CreatedAt = time.Now()
	_, err := r.db.HSetNX(context.Background(), "app:request", requestID, request).Result()
	if err != nil {
		return nil, fmt.Errorf("create: redis error: %w", err)
	}

	return request, nil
}

func (r *CacheRepository) GetAll() ([]*domain.Request, error) {

	res, err := r.db.HGetAll(context.Background(), "app:request").Result()
	if err != nil {
		return nil, fmt.Errorf("error getting requests, err: %v", err)
	}
	requests := []*domain.Request{}
	for _, stringValue := range res {
		b := &domain.Request{}
		err := json.Unmarshal([]byte(stringValue), b)
		if err != nil {
			return nil, fmt.Errorf("error decoding result, err: %v", err)
		}
		requests = append(requests, b)
	}

	return requests, nil
}

