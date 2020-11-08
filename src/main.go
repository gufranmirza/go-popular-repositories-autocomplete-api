package main

import (
	"github.com/gufranmirza/go-popular-repositories-autocomplete-api/cmd"
)

// @title Go Popular API Documentation
// @version 2.0
// @description Go Popular API Documentation

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /go-popular/v1
// @query.collection.format multi

func main() {
	cmd.Execute()
}
