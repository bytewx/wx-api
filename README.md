# wx-api

Лёгкий HTTP-сервис на Go поверх [Open-Meteo API](https://open-meteo.com/): отдаёт погоду в JSON и умеет рисовать HTML-графики (go-echarts) по температуре.

## Эндпоинты

| Метод | Путь                      | Описание                                                                 |
|-------|---------------------------|---------------------------------------------------------------------------|
| GET   | `/getWeather`              | Погода в JSON. Параметры: `weatherType` (`current`/`hourly`/`daily`), `fields` (список полей через запятую, по умолчанию — набор из `internal/config/weather.go`) |
| GET   | `/chartHourlyTemperature`  | HTML-график часовой температуры                                          |
| GET   | `/chartDailyTemperature`   | HTML-график среднесуточной температуры                                   |

## Запуск

### Локально

```bash
make run
```

По умолчанию читает конфиг из `config/config.yaml` (переопределяется переменной `CONFIG_PATH`). Redis не обязателен — при недоступности кэш просто отключается, сервис продолжает работать без него.

### Docker Compose

```bash
docker-compose up --build
```

Поднимает `wx-api` (порт 8080) и `redis` (порт 6379). Ответы `/getWeather`-семейства эндпоинтов кэшируются в Redis по ключу `weatherType+fields` с TTL из `redis.ttl_seconds` в конфиге.

## Конфигурация (`config/config.yaml`)

```yaml
server:
  host: "localhost"
  port: 8080

redis:
  addr: "redis:6379"   # для локального запуска без docker используйте "localhost:6379"
  password: ""
  db: 0
  ttl_seconds: 300
```

## Разработка

```bash
make build   # go generate + сборка бинаря в bin/
make test    # go test ./...
make lint    # golangci-lint
make tidy    # go mod tidy
```

## TODO

- [ ] Больше графиков для аналитики (осадки, ветер, давление, minutely15/pressure_levels)
- [ ] Тесты (`internal/api`, `internal/httpapi`, `internal/cache`)
