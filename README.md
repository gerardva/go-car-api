# Car API

This is a simple API for managing Car objects in a database, written in GO using Gin and GORM.

## Running

Run the API with the docker compose command:

```
docker-compose up --build
```

This will start a postgres instance and start the API on port `8080`. An environment file `app.env` is required in the root in the following format:

```
POSTGRES_HOST=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password123
POSTGRES_DB=cars
POSTGRES_PORT=5432

PORT=8080
```