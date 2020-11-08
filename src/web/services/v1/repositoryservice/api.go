package repositoryservice

import (
	"errors"
	"net/http"
)

// RepositoryService interface ...
type RepositoryService interface {
	Search(w http.ResponseWriter, r *http.Request)
}

// The list of error types presented to the end user as error message.
var (
	ErrIncompleteDetails = errors.New("Incorrect details provided, please provice correct details")
	FailedToSearchRepo   = "Failed-To-Search-Repos"
)
