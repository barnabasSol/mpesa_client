package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/barnabasSol/mpesa_client/internals/modules/auth"
	"github.com/joho/godotenv"
)

func main() {
	//not part of the go required packages but i needed it to load the consumer key and secret
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env variables")
	}

	//better built-in package for logging
	logger := slog.New(slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			AddSource: true,
		},
	))

	//client config, something similar to db connection pool but for request lifetime
	client := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    time.Second * 30,
			DisableCompression: false,
		},
	}

	authClient := auth.NewClient(client, logger)

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	result, errResp, err := authClient.GetAccessToken(consumerKey, consumerSecret)
	_ = result
	if errResp != nil {
		log.Println(errResp)
	}
	if err != nil {
		log.Fatal("omg??", err)
	}
}
