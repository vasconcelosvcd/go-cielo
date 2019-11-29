# Go-Cielo

This lib contains an client for Cielo REST API


#### Requires previous knowledge of:
- The cielo documentation
- How the api models works


### TODOS:
- [ ] Tests!!
- [ ] Better constant types
- [ ] Better constant types for currency
- [ ] One model per request/response
- [ ] One request/response per model
  
### Installation
```sh
go get github.com/vasconcelosvcd/go-cielo
```

### New Client
```golang
CieloClient, err := cielo.NewClient(os.Getenv("CIELO_MERCHANTID"), os.Getenv("CIELO_MERCHANTKEY"), cielo.SandboxEnvironment)
if err != nil {
  log.Fatal(err)
}
CieloClient.Log = os.Stdout //your current log, file or memory
```

### Tokenize Card
```golang
cc := cielo.CreditCard{
  CardNumber:"172 1263 1867 7316",
  CustomerName:"John Doe Jr",
  Holder: "John Doe" ,
  ExpirationDate:"22/2022",
  SecurityCode:"123",
  SaveCard:true,
  Brand:"VISA"
}

storedCC, err := CieloClient.CreateTokenizeCard(&cc)
if err != nil {
  log.fatal(err);
}

```

### Sale (Using an tokenized card)

```golang
sale := cielo.Sale{
  MerchantOrderID: "MY-CUSTOM-ID",
  Customer: &cielo.Customer{
    Email: "CUSTOMER-EMAIL",
    Name:  "CUSTOMER-NAME",
  },
  Payment: &cielo.Payment{
    Installments:     1,
    Type:             "CreditCard",
    Amount:           400 //Credits in cents,
    SoftDescriptor:   "MY DESCRIPTION",
    Capture:          true,
    ServiceTaxAmount: 0,
    CreditCard:       &cielo.CreditCard{CardToken: "CARD-TOKEN", SecurityCode: CCV}, 
  },
}

resp, err := CieloClient.Authorization(&sale)
if err != nil {
	log.error(err);
}
```
