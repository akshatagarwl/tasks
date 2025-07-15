```bash
sqlc generate
```

```bash
docker run --name tpg \
    -e POSTGRES_DB=tdb \
    -e POSTGRES_USER=tuser \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -p 5432:5432 \
    -v tpgdata:/var/lib/postgresql/data \
    -d postgres
```

```bash
atlas schema apply \
    --url "postgres://tuser:mysecretpassword@localhost:5432/tdb?sslmode=disable" \
    --dev-url "docker://postgres" \
    --to "file://schema.sql"
```

```bash
DB_HOST=localhost DB_PORT=5432 DB_USER=tuser DB_PASSWORD=mysecretpassword DB_NAME=tdb go run main.go
```