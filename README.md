# A simple microservice in Go

This service allows you to store and receive data in the following format.

```json
{
"external_id": "<uuid>",
"name": "some name",
"email": "email@email.com",
"date_of_birth": "2020-01-01T12:12:34+00:00"
}
```

Endpoints:

- `GET /{uuid}`
- `POST /save`

## Requirements

The project was developed with version 1.24.4 of Go.

It can also be run using Docker.

## Build

### Local

```shell
make build
```

### Docker

```shell
make build-image
```

## Usage

Start the server with

```shell
make run
```

or

```shell
docker compose up 
```

The server listens on <http://localhost:8090/>.
