package main

import (
	"log"
	"os"

	mpesa "github.com/barnabasSol/mpesa_client/pkg/lib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env variables")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	mpesaClient := mpesa.New(consumerKey, consumerSecret, nil)

}

// result, err := mpesaClient.C2B.RegisterURL(
// 	c2b.RegisterURLDto{
// 		ShortCode:       "101010",
// 		ResponseType:    "Completed",
// 		CommandID:       "RegisterURL",
// 		ConfirmationURL: "http://mydomain.com/c2b/confirmation",
// 		ValidationURL:   "http://mydomai.com/c2b/validation",
// 	},
// 	t.AccessToken,
// )
