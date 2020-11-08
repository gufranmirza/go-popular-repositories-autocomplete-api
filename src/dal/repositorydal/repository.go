package repositorydal

import (
	"context"
	"fmt"
	"time"

	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/connection"
	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/dbmodels"

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

// Create creates a new account.
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
