package repositorydal

import (
	"context"
	"fmt"
	"time"

	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/connection"
	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/dbmodels"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/spf13/viper"
)

type repo struct {
	db connection.MongoStore
}

// NewRepositoryDal ...
func NewRepositoryDal() RepositoryDal {
	return &repo{
		db: connection.NewMongoStore(),
	}
}

// Create creates a new repo into db.
func (r *repo) Create(repository *dbmodels.Repository) error {
	rc := r.db.Database().Collection(viper.GetString("db.repository_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	_, err := rc.InsertOne(ctx, repository)
	if err != nil {
		return fmt.Errorf("Failed to create repo with error %v", err)
	}

	return nil
}

// Search looks for repos matching with query.
func (r *repo) Search(query string) ([]dbmodels.Repository, error) {
	rc := r.db.Database().Collection(viper.GetString("db.repository_collection"))
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(viper.GetInt("db.query_timeout_in_sec"))*time.Second,
	)
	defer cancel()

	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"name", primitive.Regex{Pattern: query, Options: "i"}}},
			bson.D{{"full_name", primitive.Regex{Pattern: query, Options: "i"}}},
		}},
	}

	filterCursor, err := rc.Find(ctx, filter)
	if err != nil {
		return []dbmodels.Repository{}, fmt.Errorf("Failed to search repos with error %v", err)
	}

	var repos []dbmodels.Repository
	if err = filterCursor.All(ctx, &repos); err != nil {
		return []dbmodels.Repository{}, fmt.Errorf("Failed to search repos with error %v", err)
	}

	return repos, nil
}
