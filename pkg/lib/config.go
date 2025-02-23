package mpesa

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func defaultClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    time.Second * 30,
			DisableCompression: false,
		},
	}
	return client
}

func loggerConfig() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
		},
	))
	return logger
}
