package nonrelational

import (
	"context"
	"fmt"
	"gb-ui-core/config"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func BuildMongoDB(ctn di.Container) (interface{}, error) {
	cfg := ctn.Get("config").(*config.Config)

	return mongo.Connect(context.Background(), &options.ClientOptions{
		Auth: &options.Credential{
			Username: cfg.StorageConfig.NonRelational.Mongo.User,
			Password: cfg.StorageConfig.NonRelational.Mongo.Password,
		},
	},
		options.Client().
			ApplyURI(fmt.Sprintf("mongodb://%s:%s",
				cfg.StorageConfig.NonRelational.Mongo.Host,
				cfg.StorageConfig.NonRelational.Mongo.Port)))

}
