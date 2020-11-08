package repositoryservice

import (
	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/dal/repositorydal"
	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/logging"
)

type reposervice struct {
	logger        logging.Logger
	repositorydal repositorydal.RepositoryDal
}

// NewRepositoryService returns service impl
func NewRepositoryService() RepositoryService {
	return &reposervice{
		logger:        logging.NewLogger(),
		repositorydal: repositorydal.NewRepositoryDal(),
	}
}
