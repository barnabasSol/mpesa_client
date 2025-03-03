package main

import (
	"log"
	"os"

	"github.com/barnabasSol/mpesa_client/internals/modules/stkpush"
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
	// x, _, _ := mpesaClient.Auth.GetAccessToken()
	// log.Print(x.AccessToken)

	res, err := mpesaClient.STKPush.SendSTKPushRequest(stkpush.STKPushRequest{
		BusinessShortCode: "2060",
		TransactionType:   "CustomerPayBillOnline",
		Amount:            20,
		PartyA:            "251700404709",
		PartyB:            "2060",
		PhoneNumber:       "251700404709",
		AccountReference:  "Partner Unique ID",
		TransactionDesc:   "Payment Reason",
		ReferenceData: []stkpush.ReferenceData{
			{
				Key:   "ThirdPartyReference",
				Value: "Ref-12345",
			},
		},
	}, accessToken)

	if err != nil {
		log.Println(res)
		return
	}
	log.Println(res)

}
