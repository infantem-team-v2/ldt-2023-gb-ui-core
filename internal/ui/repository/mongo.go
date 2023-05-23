package repository

import (
	"context"
	"gb-ui-core/internal/ui/model"
	"gb-ui-core/pkg/terrors"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Client
}

func BuildMongoRepository(ctn di.Container) (interface{}, error) {
	return &MongoRepository{
		db: ctn.Get("mongo").(*mongo.Client),
	}, nil
}

func (m *MongoRepository) GetActiveCalculatorElements() ([]*model.UiInputElementUnitDAO, error) {
	filter := bson.D{
		{"active", true},
	}
	data := make([]*model.UiInputElementUnitDAO, 0, 20)

	cursor, err := m.db.
		Database(model.UIMongoDB).
		Collection(model.CalculatorCollection).
		Find(context.Background(), filter)
	if err != nil {
		return nil, terrors.Raise(err, 300008)
	}

	for cursor.Next(context.Background()) {
		var element model.UiInputElementUnitDAO
		err = cursor.Decode(&element)
		if err != nil {
			return nil, terrors.Raise(err, 300009)
		}
		data = append(data, &element)
	}
	if err := cursor.Close(context.Background()); err != nil {
		return nil, terrors.Raise(err, 300010)
	}
	return data, nil
}
