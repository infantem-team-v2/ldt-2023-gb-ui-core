package repository

import (
	"context"
	calcModel "gb-ui-core/internal/calculator/model"
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

func (m *MongoRepository) UpdateActiveElements(params *calcModel.SetActiveForElementRequest) error {
	for _, unit := range params.Elements {
		_, err := m.db.
			Database(model.UIMongoDB).
			Collection(model.CalculatorCollection).
			UpdateOne(context.Background(),
				bson.D{{
					"field_id", unit.FieldId,
				}},
				bson.D{{
					"$set", bson.D{{
						"active", unit.Active,
					}},
				}},
			)
		return terrors.Raise(err, 300011)
	}

	return nil
}

func (m *MongoRepository) GetActiveElementsByCategory(doAdmin bool) ([]*model.UiInputCategoryDAO, error) {
	groupPipeline := bson.D{{
		"$group", bson.D{
			{"_id", "$category"},
			{"elements", bson.D{{
				"$push", bson.D{
					{"field", "$field"},
					{"type", "$type"},
					{"active", "$active"},
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
	matchPipeline := bson.D{{
		"$match", bson.D{{
			"active", true,
		}},
	}}
	var pipeline mongo.Pipeline
	if doAdmin {
		pipeline = mongo.Pipeline{
			groupPipeline, projectPipeline,
		}
	} else {
		pipeline = mongo.Pipeline{
			matchPipeline, groupPipeline, projectPipeline,
		}
	}

	cursor, err := m.db.
		Database(model.UIMongoDB).
		Collection(model.CalculatorCollection).
		Aggregate(context.Background(), pipeline)

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
