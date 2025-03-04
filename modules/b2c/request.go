package b2c

type B2CRequest struct {
	InitiatorsName     string `json:"InitiatorsName"`
	SecurityCredential string `json:"SecurityCredential"`
	Occassion 	       string `json:"Occassion"`
	CommandID 	       string `json:"CommandID"`
	PartyA 		       string `json:"PartyA"`
	PartyB 		       string `json:"PartyB"`
	Remarks 	       string `json:"Remarks"`
	Amount 		       string `json:"Amount"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL 		   string `json:"ResultURL"`
}