# RSS Aggregator

## Use SQLC to generate Golang Code from SQL
- install SQLC
```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```
- run `sqlc generate` to generate the code for sql query files

## Database Migration With Goose
- install Goose
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- create database `rss_aggregator` in postgres database
- this command need to run in `migrations/schema` folder
``` To Create Tables
goose postgres postgres://<dbUsername>:<dbPassword>@localhost:3000/rss_aggregator up
```
``` To Drop Tables or columns based on the last database migration action
goose postgres postgres://<dbUsername>:<dbPassword>@localhost:3000/rss_aggregator down
```
``` To Drop All Tables
goose postgres postgres://<dbUsername>:<dbPassword>@localhost:3000/rss_aggregator down-to 0
```

## Install TailwindCSS
```
npm i
```

## Hot Reload & Browser Refresh
- Install Air
```
go install github.com/air-verse/air@latest
```
- If your are using html templates for ui and wanted to auto refresh the browser eveytime app rebuild, add `dev` in `APP_ENV` in `.env` file
```
APP_ENV="dev"
```
- Run Air
```
air
```

- NOTE: if u want to exclude some folder from hot reload list, add those folder name in `exclude_dir` list or `exclude_file` for file list in `.air.toml` file
```
exclude_dir = ["tmp", "vendor", "testdata", "node_modules", "api_docs", "migrations", "unused_codes"]
exclude_file = ["static_assets/build/style.css"]
```