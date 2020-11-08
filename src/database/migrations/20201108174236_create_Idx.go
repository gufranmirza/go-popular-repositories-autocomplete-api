package migrations

import (
	"context"
	"time"

	"github.com/spf13/viper"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	mod := []mongo.IndexModel{
		{
			Keys:    bson.M{"id": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"name": 1},
		},
		{
			Keys: bson.M{"full_name": 1},
		},
	}

	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(viper.GetString("db.repository_collection")).Indexes().CreateMany(ctx, mod)
		return err
	}, func(db *mongo.Database) error { //Down
		_, err := db.Collection(viper.GetString("db.repository_collection")).Indexes().DropOne(ctx, "id_1")
		if err != nil {
			return err
		}
		_, err = db.Collection(viper.GetString("db.repository_collection")).Indexes().DropOne(ctx, "name_1")
		if err != nil {
			return err
		}
		_, err = db.Collection(viper.GetString("db.repository_collection")).Indexes().DropOne(ctx, "full_name_1")
		return err
	})
}
