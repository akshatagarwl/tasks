# Tasks API

- API Docs: https://tasks-api.akshat.dev/swagger/index.html
- Docker Images: https://hub.docker.com/r/akshataga/tasks
- ER diagram: https://gh.atlasgo.cloud/explore/bc8dd5b0

## Example cURLs

Create task
```bash
curl -s -X POST "https://tasks-api.akshat.dev/task" \
  -H "accept: application/json" -H "Content-Type: application/json" \
  -d '{"title": "Update README", "description": "Add descriptive documentation", "status": "TODO"}'
```

Get task
```bash
curl -s -X GET "https://tasks-api.akshat.dev/task/<UUID>" -H "accept: application/json"
```

Get tasks by statuses
```bash
curl -s -X GET "https://tasks-api.akshat.dev/task?statuses=TODO,IN_PROGRESS" -H "accept: application/json"
```

Update task
```bash
curl -s -X PUT "https://tasks-api.akshat.dev/task/<UUID>" \
  -H "accept: application/json" -H "Content-Type: application/json" \
  -d '{"title": "Updated Task", "description": "Updated description", "status": "IN_PROGRESS"}'
```

Delete task
```bash
curl -s -X DELETE "https://tasks-api.akshat.dev/task/<UUID>" -H "accept: application/json"
```

## Deployment
- Koyeb for API
- Neon for DB

## Stack
- Go + Fiber
- Postgres
- SQLC (queries → type-safe Go)
- Atlas (migrations)

# Quick Start

## Docker Compose (Recommended)

You will need an Atlas API token to run this. To create a token, see [Atlas bots](https://atlasgo.io/cloud/bots) and [Getting Started](https://atlasgo.io/cloud/getting-started).

```bash
ATLAS_TOKEN=<token> docker compose up --build
```

## Local Setup
```bash
# generate Swagger docs
go run github.com/swaggo/swag/cmd/swag@latest init
```

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

## Considerations during development

- Pagination: Uses an offset based pagination
- Task Status as Enum in DB: This makes the schema migrations a bit more complex but suffices for now.
