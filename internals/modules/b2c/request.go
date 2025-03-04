package b2c

type B2CRequest struct {
	InitiatorsName  string    `json:"InitiatorsName"`
	Initiator 	    Initiator `json:"Initiator"`
	Occassion 	    string	  `json:"Occassion"`
	CommandID 	    string	  `json:"CommandID"`
	PartyA 		    string	  `json:"PartyA"`
	PartyB 		    string	  `json:"PartyB"`
	Remarks 	    string	  `json:"Remarks"`
	Amount 		    string	  `json:"Amount"`
	QueueTimeOutURL string	  `json:"QueueTimeOutURL"`
	ResultURL 		string	  `json:"ResultURL"`
}

type Initiator struct {
	IdentifierType     int    `json:"IdentifierType"`
	Identifier         string `json:"Identifier"`
	SecurityCredential string `json:"SecurityCredential"`
	SecretKey          string `json:"SecretKey"`
}