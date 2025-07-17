# Swagger API Documentation

This project now includes comprehensive Swagger/OpenAPI documentation for the Task Management API.

## Features Added

- **Swagger UI**: Interactive API documentation available at `/docs/swagger/index.html`
- **OpenAPI 2.0 Specification**: Complete API specification in JSON and YAML formats
- **Request/Response Examples**: All endpoints include example data
- **Parameter Documentation**: Query parameters, path parameters, and request bodies are fully documented

## API Endpoints

### Tasks
- `GET /task` - Get tasks with optional filtering and pagination
- `POST /task` - Create a new task
- `GET /task/{id}` - Get a task by ID
- `PUT /task/{id}` - Update a task
- `DELETE /task/{id}` - Delete a task

### Health Check
- `GET /livez` - Health check endpoint

## Task Status Values
- `PENDING` - Task is pending
- `IN_PROGRESS` - Task is in progress
- `COMPLETED` - Task is completed

## Usage

1. Start the application:
   ```bash
   go run main.go
   ```

2. Access the Swagger UI at: `http://localhost:8080/docs/swagger/index.html`

3. Use the interactive interface to:
   - View all available endpoints
   - Test API calls directly from the browser
   - View request/response schemas
   - See example data for all models

## Environment Variables

Ensure these environment variables are set:
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `SERVER_PORT` - Server port (default: 8080)

## Generated Files

- `docs/docs.go` - Generated Go documentation
- `docs/swagger.json` - OpenAPI specification in JSON format
- `docs/swagger.yaml` - OpenAPI specification in YAML format

## Regenerating Documentation

To regenerate the Swagger documentation after making changes:

```bash
swag init
```

This will update the documentation files in the `docs/` directory.