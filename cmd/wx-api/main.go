package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bytewx/wx-api/internal/cache"
	"github.com/bytewx/wx-api/internal/config"
	"github.com/bytewx/wx-api/internal/httpapi"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config/config.yaml"
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("fatal error: loading config: %v", err)
	}

	cacheClient, err := cache.NewClient(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)
	if err != nil {
		log.Printf("warning: redis unavailable at %s, running without caching: %v", cfg.Redis.Addr, err)
		cacheClient = nil
	} else {
		log.Printf("info: connected to redis at %s", cfg.Redis.Addr)
	}

	handler := httpapi.NewHandler(cacheClient, cfg.Redis.TTLSeconds)

	router := http.NewServeMux()

	router.HandleFunc("GET /getWeather", handler.HandleGetWeather)
	router.HandleFunc("GET /chartHourlyTemperature", handler.HandleChartHourlyTemperature)
	router.HandleFunc("GET /chartDailyTemperature", handler.HandleChartDailyTemperature)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	go func() {
		log.Printf("Listening on %s:%d", cfg.Server.Host, cfg.Server.Port)
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("fatal error: server listen failed: %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("info: shutdown signal received, starting graceful shutdown...")

	shutdownCtx, stop := context.WithTimeout(context.Background(), 5*time.Second)
	defer stop()

	if err = server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("fatal error: server shutdown failed: %v", err)
	}

	log.Println("info: graceful shutdown complete")
}
