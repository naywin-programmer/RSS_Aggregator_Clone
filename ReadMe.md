# RSS Aggregator


## Use SQLC to generate Golang Code from SQL
- install SQLC
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
- create `sqlc.yaml' file
- run `sqlc generate` to generate the code for sql query files


## Database Migration With Goose
- install Goose
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- create database called "rss_aggregator"
- the command need to run in the same place that have `.sql` files
```
goose postgres postgres://dbUsername:dbPassword@localhost:5432/rss_aggregator up
goose postgres postgres://dbUsername:dbPassword@localhost:5432/rss_aggregator down
```
