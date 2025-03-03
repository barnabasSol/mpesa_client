package main

import (
	"log"
	"os"

	"github.com/barnabasSol/mpesa_client/internals/modules/c2b"
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
	accessToken := os.Getenv("ACCESS_TOKEN")
	_ = accessToken

	mpesaClient := mpesa.New(
		mpesa.Sandbox,
		nil,
		consumerKey,
		consumerSecret,
	)
	_ = mpesaClient
	res, err := mpesaClient.C2B.ProcessPayment(c2b.PaymentRequest{}, accessToken)
	if err != nil {
		log.Println(res)
		return
	}

}
