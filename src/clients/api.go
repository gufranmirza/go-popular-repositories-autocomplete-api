package clients

import "github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/dbmodels"

// Clients implents the clients interface
type Clients interface {
	GetRepositories() ([]dbmodels.Repository, error)
}
