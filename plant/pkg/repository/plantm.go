package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	domain "github.com/garhus2020/ESIhw2/plant/pkg/domain"
)

type PlantmRepository struct {
	db *mongo.Client
}

func NewPlantmRepository(db *mongo.Client) *PlantmRepository {
	return &PlantmRepository{
		db: db,
	}
}

func (r *PlantmRepository) Create(plantm *domain.Plantm) (*domain.Plantm, error) {
	collection := r.db.Database("local").Collection("plantm")

	plantm.CreatedAt = time.Now()
	res, err := collection.InsertOne(context.Background(), plantm)
	if err != nil {
		return nil, fmt.Errorf("error inserting plantm %v, err: %v", plantm, err)
	}
	plantm.ID = res.InsertedID

	return plantm, nil
}

func (r *PlantmRepository) GetAll() ([]*domain.Plantm, error) {
	collection := r.db.Database("local").Collection("plantm")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting plantms, err: %v", err)
	}
	plantms := []*domain.Plantm{}
	for cursor.Next(context.Background()) {
		b := &domain.Plantm{}
		err := cursor.Decode(&b)
		if err != nil {
			return nil, fmt.Errorf("error decoding result, err: %v", err)
		}
		plantms = append(plantms, b)
	}
	err = cursor.Close(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting plantms, err: %v", err)
	}

	return plantms, nil
}

func (r *PlantmRepository) GetOne(ident string) (string, error) {
	collection := r.db.Database("local").Collection("plantm")
	//cursor, err := collection.Find(context.Background(), bson.D{})
	// if err != nil {
	// 	return "nil", fmt.Errorf("error getting plantms, err: %v", err)
	// }
	plantm := &domain.Plantm{}
	filter := bson.M{"ident": ident}
	err := collection.FindOne(context.Background(), filter).Decode(&plantm)
	if err != nil {
		return "nil", fmt.Errorf("error getting plantms, err: %v", err)
	}

	// err = cursor.Close(context.Background())
	// if err != nil {
	// 	return "nil", fmt.Errorf("error getting plantms, err: %v", err)
	// }

	return plantm.Price, nil
}
