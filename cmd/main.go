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

	mpesaClient := mpesa.New(
		"https://apisandbox.safaricom.et",
		consumerKey,
		consumerSecret,
		nil,
	)
	mpesaClient.Auth.GetAccessToken()
	//test
	mpesaClient.STKPush.SendSTKPushRequest("accesstoken", "1234", "paybill", "100", "098765", "1234", "DATA", "desc")

}
