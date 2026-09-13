# wx-api

A lightweight Go HTTP service on top of [Open-Meteo API](https://open-meteo.com/): serves weather data as JSON and renders HTML temperature charts (go-echarts).

## Endpoints

| Method | Path                       | Description                                                                                   |
|--------|----------------------------|-------------------------------------------------------------------------------------------------|
| GET    | `/getWeather`              | Weather as JSON. Params: `weatherType` (`current`/`hourly`/`daily`), `fields` (comma-separated list; defaults to the set in `internal/config/weather.go`) |
| GET    | `/chartHourlyTemperature`  | HTML chart of hourly temperature                                                                |
| GET    | `/chartDailyTemperature`   | HTML chart of average daily temperature                                                         |

## Running

### Locally

```bash
make run
```

By default it reads config from `config/config.yaml` (override with the `CONFIG_PATH` env variable). Redis is optional — if it's unavailable, caching is simply disabled and the service keeps running without it.

### Docker Compose

```bash
docker-compose up --build
```

Spins up `wx-api` (port 8080) and `redis` (port 6379). Responses from the `/getWeather` family of endpoints are cached in Redis under a `weatherType+fields` key, with TTL taken from `redis.ttl_seconds` in the config.

## Configuration (`config/config.yaml`)

```yaml
server:
  host: "localhost"
  port: 8080

redis:
  addr: "redis:6379"
  password: ""
  db: 0
  ttl_seconds: 300
```

## Development

```bash
make build   # go generate + build binary into bin/
make test    # go test ./...
make lint    # golangci-lint
make tidy    # go mod tidy
```
