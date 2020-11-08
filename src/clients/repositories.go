package clients

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/database/dbmodels"
)

type clientsImpl struct{}

// NewClient returns the client interface
func NewClient() Clients {
	return &clientsImpl{}
}

// JSONData contains the GitHub API response
type JSONData struct {
	Count int `json:"total_count"`
	Items []dbmodels.Repository
}

// GetRepositories returns the top 100 Golang repositories
func (c *clientsImpl) GetRepositories() ([]dbmodels.Repository, error) {
	res, err := http.Get("https://api.github.com/search/repositories?q=stars:>=5000+language:go&sort=stars&order=desc&per_page=300")
	if err != nil {
		return []dbmodels.Repository{}, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return []dbmodels.Repository{}, err
	}
	if res.StatusCode != http.StatusOK {
		return []dbmodels.Repository{}, fmt.Errorf("Unexpected status code %v", res.StatusCode)
	}
	data := JSONData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []dbmodels.Repository{}, fmt.Errorf("Failed to unmarshal with error %v", err)
	}

	return data.Items, nil
}
