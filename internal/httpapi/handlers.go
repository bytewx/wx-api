package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bytewx/wx-api/internal/api"
	"github.com/bytewx/wx-api/internal/cache"
	"github.com/bytewx/wx-api/internal/config"
	"github.com/bytewx/wx-api/internal/utils"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Handler bundles the dependencies needed by the HTTP handlers. Cache may be nil,
// in which case requests always hit the upstream weather API directly.
type Handler struct {
	Cache      *cache.Client
	TTLSeconds int
}

func NewHandler(cacheClient *cache.Client, ttlSeconds int) *Handler {
	return &Handler{Cache: cacheClient, TTLSeconds: ttlSeconds}
}

func (h *Handler) HandleGetWeather(w http.ResponseWriter, r *http.Request) {
	const operation = "wx-api.internal.httpapi.handlers.HandleWeather"

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	weatherType := r.URL.Query().Get("weatherType")
	fields := r.URL.Query().Get("fields")

	cfg, ok := config.WeatherTypes[config.WeatherType(weatherType)]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	response, err := api.GetWeather(h.Cache, h.TTLSeconds, fields, config.WeatherType(cfg.ParamName))
	if err != nil {
		fmt.Printf("%s: error: get weather: %s\n", operation, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("%s: error: encode response: %s\n", operation, err)
	}
}

func (h *Handler) HandleChartHourlyTemperature(w http.ResponseWriter, r *http.Request) {
	const operation = "wx-api.internal.httpapi.handlers.HandleChartHourlyTemperature"

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	response, err := api.GetWeather(h.Cache, h.TTLSeconds, "temperature_2m", config.Hourly)
	if err != nil {
		fmt.Printf("%s: unexpected error happened trying to fetch hourly weather: %s\n",
			operation,
			err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if response.HourlyWeather.Temperature2M == nil || response.HourlyWeather.Time == nil {
		http.Error(w, "No temperature data available", http.StatusInternalServerError)
		return
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Hourly Temperature Forecast",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Time",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Temperature (°C)",
		}),
	)

	line.SetXAxis(response.HourlyWeather.Time).
		AddSeries("Temperature 2M", utils.FloatsToLineData(response.HourlyWeather.Temperature2M))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = line.Render(w); err != nil {
		fmt.Printf("%s: unexpected error happened trying to render chart: %s\n",
			operation,
			err,
		)
	}
}

func (h *Handler) HandleChartDailyTemperature(w http.ResponseWriter, r *http.Request) {
	const operation = "wx-api.internal.httpapi.handlers.HandleGetDailyTemperature"

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	response, err := api.GetWeather(h.Cache, h.TTLSeconds, "temperature_2m_mean", config.Daily)
	if err != nil {
		fmt.Printf("%s: unexpected error happened trying to fetch daily weather: %s\n",
			operation,
			err,
		)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if response.DailyWeather.Temperature2MMean == nil || response.DailyWeather.Time == nil {
		http.Error(w, "No temperature data available", http.StatusInternalServerError)
		return
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Daily Temperature Forecast",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Time",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Temperature (°C)",
		}),
	)

	line.SetXAxis(response.DailyWeather.Time).
		AddSeries("Temperature Mean", utils.FloatsToLineData(response.DailyWeather.Temperature2MMean))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = line.Render(w); err != nil {
		fmt.Printf("%s: unexpected error happened trying to render chart: %s\n",
			operation,
			err,
		)
	}
}
