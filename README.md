# Tasks API

## Stack
- Go + Fiber
- Postgres
- SQLC (queries → type-safe Go)
- Atlas (migrations)

## Local Setup
```bash
# generate typed queries
sqlc generate

# start Postgres
docker run --name tpg -e POSTGRES_DB=tdb -e POSTGRES_USER=tuser \
  -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 \
  -v tpgdata:/var/lib/postgresql/data -d postgres

# apply schema
atlas schema apply \
  --url "postgres://tuser:mysecretpassword@localhost:5432/tdb?sslmode=disable" \
  --dev-url "docker://postgres" --to "file://schema.sql"

# run api
DB_HOST=localhost DB_PORT=5432 DB_USER=tuser \
DB_PASSWORD=mysecretpassword DB_NAME=tdb go run main.go
```

## Docker Compose
```bash
ATLAS_TOKEN=<token> docker compose up --build
```

## Docs
- Swagger UI: `/swagger`
- ER diagram: https://gh.atlasgo.cloud/explore/bc8dd5b0

Production: https://tasks-api.akshat.dev
