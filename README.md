# go-popular-repositories-autocomplete-api
Provides Autocomplete functionality for Popular Golang Repositories 


# Running Project
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

## Testing
Running test
```
make test
```

## Mocks
Generate package mocks
```
make mock
```

# Docker
Build and run docker image
```
docker-compose up -d --build
```

## Visit
[http://localhost:8001/go-popular/v1/repository/search?q=core&limit=5](http://localhost:8001/go-popular/v1/repository/search?q=core&limit=5)

[Swagger API documetation](http://localhost:8001/go-popular/v1/swagger/index.html)
