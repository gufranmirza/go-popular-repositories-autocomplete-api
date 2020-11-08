package repositorydal

import "github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/dbmodels"

// RepositoryDal ...
type RepositoryDal interface {
	Create(repository *dbmodels.Repository) error
}
