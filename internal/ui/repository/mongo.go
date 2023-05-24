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

func (m *MongoRepository) GetActiveElementsByCategory() ([]*model.UiInputCategoryDAO, error) {
	groupPipeline := bson.D{{
		"$group", bson.D{
			{"_id", "$category"},
			{"elements", bson.D{{
				"$push", bson.D{
					{"field", "$field"},
					{"type", "$type"},
					{"field_id", "$field_id"},
					{"comment", "$comment"},
					{"options", "$options"},
				},
			},
			}},
		}}}
	projectPipeline := bson.D{{
		"$project", bson.D{{
			"elements", bson.D{{
				"$slice", bson.A{
					"$elements", 50,
				},
			}},
		}},
	}}

	cursor, err := m.db.
		Database(model.UIMongoDB).
		Collection(model.CalculatorCollection).
		Aggregate(context.Background(), mongo.Pipeline{
			groupPipeline, projectPipeline,
		})

	if err != nil {
		return nil, terrors.Raise(err, 300008)
	}

	data := make([]*model.UiInputCategoryDAO, 0, 50)

	err = cursor.All(context.Background(), &data)

	if err != nil {
		return nil, terrors.Raise(err, 300009)
	}

	return data, nil
}
