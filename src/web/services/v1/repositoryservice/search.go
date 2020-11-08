package repositoryservice

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/web/renderers"
)

// @Summary Search repositories
// @Description It allows to search repositories
// @Tags jobs
// @Accept json
// @Produce json
// @Param q query string true "search query"
// @Param limit query string true "response limit default is 10"
// @Success 200 {object} []dbmodels.Repository{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /repository/search [GET]
func (repo *reposervice) Search(w http.ResponseWriter, r *http.Request) {
	txID := r.Header["transaction_id"][0]
	query := r.URL.Query().Get("q")
	limitQuery := r.URL.Query().Get("limit")
	limit := 10

	if len(limitQuery) > 0 {
		i, err := strconv.Atoi(limitQuery)
		if err == nil {
			limit = i
		}
	}

	if len(query) == 0 {
		repo.logger.Error(txID, FailedToSearchRepo).Infof("Invalid query %s", query)
		render.Render(w, r, renderers.ErrorInvalidRequest(errors.New("Query can not be empty")))
		return
	}

	resp, err := repo.repositorydal.Search(query)
	if err != nil {
		repo.logger.Error(txID, FailedToSearchRepo).Error(err)
		render.Render(w, r, renderers.ErrorNotFound(err))
		return
	}

	if len(resp) > limit {
		resp = resp[:limit]
	}

	render.Respond(w, r, resp)
	return
}
