The **Minimal MPESA Client Library** is a lightweight SDK designed to simplify interactions with MPESA services. It provides an intuitive interface for developers to integrate MPESA functionality, such as authentication and payment processing, into their applications. 
#### Getting Started

To initialize the MPESA client first get the library then use the following code snippet:

```bash
go get github.com/barnabasSol/mpesa_client
```

```go
mpesaClient := mpesa.New(
    mpesa.Sandbox,        // Environment: mpesa.Sandbox or mpesa.Prod
    nil,                  // HTTP client: nil for default or provide your custom client
    consumerKey,          // Consumer Key from the MPESA Developer Portal
    consumerSecret,       // Consumer Secret from the MPESA Developer Portal
)
```

- **Environment**: Choose between `mpesa.Sandbox` for development or `mpesa.Production` for live transactions.
- **HTTP Client**: Use `nil` to rely on the default HTTP client or supply your own custom configuration.
- **Consumer Key & Secret**: Obtain these credentials from the MPESA Developer Portal.

#### Authentication Example

To retrieve an access token, use the following code:

```go
result, respErr, err := mpesaClient.Auth.GetAccessToken()
log.Print(result.AccessToken)
```

This snippet authenticates with the MPESA API and logs the generated access token.

---

### Making an STK Push Request

The library also simplifies initiating an STK Push (Lipa na M-Pesa Online Payment). Here's an example:

```go
res, err := mpesaClient.STKPush.SendRequest(stkpush.Request{
    BusinessShortCode: "4646",               
    TransactionType:   "CustomerPayBillOnline", 
    Amount:            20,                  
    PartyA:            "251700404709",      
    PartyB:            "4646",             
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
    log.Println("Error:", err)
    return
}
log.Println("Response:", res)
```

#### Key Parameters:
- **BusinessShortCode**: Your business identifier.
- **TransactionType**: Type of transaction (e.g., `CustomerPayBillOnline`).
- **Amount**: Payment amount.
- **PartyA/PartyB**: Customer and business details.
- **ReferenceData**: Optional metadata for additional context.

### Registering a URL for C2B

In addition to STK Push, the library also supports registering URLs for C2B (Customer to Business) transactions. This involves specifying callback URLs for transaction validation and confirmation.

#### Registering a URL

To register a URL for C2B, you can use the `RegisterURL` function:

```go
registerDto := RegisterURLDto{
    ShortCode:       "802000",               
    ResponseType:    "Completed",           
    CommandID:       "RegisterURL",         
    ConfirmationURL:  "https://www.myservice:8080/confirmation", 
    ValidationURL:   "https://www.myservice:8080/validation",   
}

consumerKey := "you-consumer-key" 

response, err := RegisterURL(registerDto, consumerKey)

if err != nil {
    log.Println("Error registering URL:", err)
    return
}

log.Println("Registration Response:", response)
```

### Making B2C request
This function allows you to do M-Pesa transactions from a company to a client.

```go
response, err := mpesaClient.B2C.SendRequest(b2c.Request{
        InitiatorsName:      "YourInitiatorsName",
        SecurityCredentials: "iSHJEgQYt3xidNVJ7lbXZqRXUlBqpM",
        Occassion:           "Disbursement",
        CommandID:           "BusinessPayment",
        PartyA:              "101010",
        PartyB:              "251700100100", 
        Remarks:             "test",
        Amount:              "1000.00",
        QueueTimeOutURL:     "https://example.com/timeout",
        ResultURL:           "https://example.com/result",
    }, accessToken)

 if err != nil {
    log.Printf("Error sending B2C request: %v\n", err)
    return
}
```
