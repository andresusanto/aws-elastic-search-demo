# Ingester &middot; [![Build](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/build-ingester.yml/badge.svg)](https://github.com/andresusanto/aws-elastic-search-demo/actions/workflows/build-ingester.yml)

Ingester is responsible for ingesting events into ElasticSearch. In short, this app:

- provides a HTTP endpoint to ingest user-click events.
- supports AWS ElasticSearch authentication using signed-requests using ECS Roles.
- is packaged using Docker and deployed on ECS.

### The Stack

1. **App:** Golang with Gin
2. **Code Standard and Quality**: Go fmt, staticcheck.io
3. **Unit Testing:** Go test
4. **Logging:** `zerolog`
5. **ES Client:** `olivere/elastic/v7`
6. **CI/CD:** GitHub Action to automatically test, build, and push images on mainline branch.

### Docker Images

```bash
docker pull ghcr.io/andresusanto/es-event-ingester:<TAG>
```

See all available tags [here](https://github.com/andresusanto/aws-elastic-search-demo/pkgs/container/es-event-ingester).

### Developing

**Requirements:**

1. Golang v1.17. It's recommended to use [gvm](https://github.com/moovweb/gvm) to switch between Go versions.
2. Docker v20.x or newer
3. [Air](https://github.com/cosmtrek/air) (automatic-reloader)
4. [Golint](https://github.com/golang/lint) and [Static Check](https://staticcheck.io/docs/install) (code linting)

**Before developing:**

```bash
# Run a local elastic search instance using Docker

$ docker run -d --name es-ingester \
    -p 9200:9200 -p 9300:9300 \
    -e "discovery.type=single-node" elasticsearch:7.12.1

# Install all required tool and dependencies:

$ go mod download
```

**When Developing:**

```bash
# Using Air (auto-reloader)

$ air


# Using go run

$ DEBUG=true DEVELOP=true go run ./cmd/server
```

**After making changes:**

```bash
# Perform code formatting (if you do not have Go fmt integration with your IDE)

$ go fmt ./...


# Perform code linting

$ golint ./... && staticcheck ./...


# Perform unit tests

$ go test ./...
```

### Building

Requirements:

1. Docker v20.x or newer

Steps:

```bash
# Run the docker build command:
$ docker build -t <NAME>:<TAG> .


# Push the built image to registry
$ docker push <NAME>:<TAG>
```

### Environment Variables

See [config.go](./internal/config/config.go) for more details.

| Environment      | Type      | Description                                                | Default Value           |
| ---------------- | --------- | ---------------------------------------------------------- | ----------------------- |
| `PORT`           | _integer_ | the port in which the app should listen to.                | `8080`                  |
| `REGION`         | _string_  | AWS Region (used to sign ElasticSearch requests)           |                         |
| `ES_ENDPOINT`    | _string_  | ElasticSearch Endpoint                                     | `http://localhost:9200` |
| `SIGN_ES_CLIENT` | _boolean_ | If true ES requests will be signed using detected ECS Role | `false`                 |
| `DEVELOP`        | _boolean_ | If true console output will be human friendly              | `false`                 |
| `DEBUG`          | _boolean_ | If true log level will be set to DEBUG                     | `false`                 |
