package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bytewx/wx-api/internal/cache"
	"github.com/bytewx/wx-api/internal/config"
	"github.com/bytewx/wx-api/internal/models"
)

// GetWeather fetches weather data for the given fields/weatherType, using cacheClient
// as a read-through cache when it is non-nil. If cacheClient is nil (e.g. Redis is
// unavailable), it falls back to always fetching from the upstream API.
func GetWeather(
	cacheClient *cache.Client,
	ttlSeconds int,
	fields string,
	weatherType config.WeatherType,
) (models.Response, error) {
	const operation = "wx-api.internal.api.handlers.GetWeather"

	cfg, ok := config.WeatherTypes[weatherType]
	if !ok {
		return models.Response{}, errors.New("weather type not found")
	}

	if fields == "" {
		fields = cfg.DefaultFields
	}

	cacheKey := fmt.Sprintf("wx-api:%s:%s", weatherType, fields)

	if cacheClient != nil {
		if response, ok := getCached(cacheClient, cacheKey); ok {
			return response, nil
		}
	}

	response, err := fetchWeather(fields, cfg.ParamName)
	if err != nil {
		return models.Response{}, err
	}

	if cacheClient != nil {
		setCached(cacheClient, cacheKey, response, ttlSeconds)
	}

	return response, nil
}

func getCached(cacheClient *cache.Client, key string) (models.Response, bool) {
	const operation = "wx-api.internal.api.handlers.getCached"

	cached, err := cacheClient.Get(context.Background(), key)
	if err != nil {
		// cache miss or redis error - not fatal, just fall through to a live fetch
		return models.Response{}, false
	}

	var response models.Response
	if err = json.Unmarshal([]byte(cached), &response); err != nil {
		fmt.Printf("%s: warning: failed to unmarshal cached response: %v\n", operation, err)
		return models.Response{}, false
	}

	return response, true
}

func setCached(cacheClient *cache.Client, key string, response models.Response, ttlSeconds int) {
	const operation = "wx-api.internal.api.handlers.setCached"

	data, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("%s: warning: failed to marshal response for cache: %v\n", operation, err)
		return
	}

	if err = cacheClient.Set(context.Background(), key, string(data), ttlSeconds); err != nil {
		fmt.Printf("%s: warning: failed to write cache entry: %v\n", operation, err)
	}
}

func fetchWeather(fields, paramName string) (models.Response, error) {
	const operation = "wx-api.internal.api.handlers.fetchWeather"

	params := url.Values{}
	params.Set("latitude", "52.52")
	params.Set("longitude", "13.41")
	params.Set(paramName, fields)

	fullURL := "https://api.open-meteo.com/v1/forecast?" + params.Encode()

	resp, err := http.Get(fullURL)
	if err != nil {
		return models.Response{}, fmt.Errorf(
			"%s: error: fetch weather: %w",
			operation,
			err,
		)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Printf("%s: error: close response body: %v\n", operation, err)
		}
	}()

	var response models.Response
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.Response{}, fmt.Errorf(
			"%s: error: decode response body: %w",
			operation,
			err,
		)
	}

	return response, err
}
