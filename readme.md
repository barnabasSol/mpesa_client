### Introduction to the Minimal MPESA Client Library

The **Minimal MPESA Client Library** is a lightweight SDK designed to simplify interactions with MPESA services. It provides an intuitive interface for developers to integrate MPESA functionality, such as authentication and payment processing, into their applications. The library supports both **sandbox** and **production** environments, making it ideal for testing and deployment.

#### Getting Started

To initialize the MPESA client, use the following code snippet:

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
