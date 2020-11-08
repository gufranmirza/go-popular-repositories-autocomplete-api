# go-popular-repositories-autocomplete-api
Provides Autocomplete functionality for Popular Golang Repositories 


# Development
## Build Project
```
make
```

## Apply Migrations
```
make migrate-up
```

## Load Data
Load Github repository data from the Github API into the database
```
make dataloader
```

## Run
```
make run
```

## Search API
[http://localhost:8001/go-popular/v1/repository/search?q=core&limit=5](http://localhost:8001/go-popular/v1/repository/search?q=core&limit=5)

[Swagger API documetation](http://localhost:8001/go-popular/v1/swagger/index.html)


