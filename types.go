package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type (
	Environment struct {
		APIUrl, APIQueryURL string
	}

	// Client represents a Cielo REST API Client
	Client struct {
		sync.Mutex
		Client      *http.Client
		MerchantId  string
		MerchantKey string
		Environment Environment
		Log         io.Writer // If user set log file name all requests will be logged there
	}

	ErrorResponse struct {
		Response *http.Response `json:"-"`
		Name     uint32         `json:"code"`
		Message  string         `json:"message"`
	}

	Sale struct {
		MerchantOrderID string    `json:",omitempty"`
		Customer        *Customer `json:",omitempty"`
		Payment         *Payment  `json:",omitempty"`
	}

	Customer struct {
		Name            string   `json:",omitempty"`
		Email           string   `json:",omitempty"`
		BirthDate       string   `json:",omitempty"`
		Identity        string   `json:",omitempty"`
		IdentityType    string   `json:",omitempty"`
		Address         *Address `json:",omitempty"`
		DeliveryAddress *Address `json:",omitempty"`
	}

	Address struct {
		Street     string `json:",omitempty"`
		Number     string `json:",omitempty"`
		Complement string `json:",omitempty"`
		ZipCode    string `json:",omitempty"`
		City       string `json:",omitempty"`
		State      string `json:",omitempty"`
		Country    string `json:",omitempty"`
	}

	Payment struct {
		ServiceTaxAmount    uint32            `json:",omitempty"`
		Installments        uint32            `json:",omitempty"`
		Interest            interface{}       `json:",omitempty"`
		Capture             bool              `json:",omitempty"`
		Authenticate        bool              `json:",omitempty"`
		Recurrent           bool              `json:",omitempty"`
		RecurrentPayment    *RecurrentPayment `json:",omitempty"`
		CreditCard          *CreditCard       `json:",omitempty"`
		DebitCard           *DebitCard        `json:",omitempty"`
		Tid                 string            `json:",omitempty"`
		ProofOfSale         string            `json:",omitempty"`
		AuthorizationCode   string            `json:",omitempty"`
		SoftDescriptor      string            `json:",omitempty"`
		ReturnURL           string            `json:",omitempty"`
		Provider            string            `json:",omitempty"`
		PaymentID           string            `json:",omitempty"`
		Type                string            `json:",omitempty"`
		Amount              uint32            `json:",omitempty"`
		ReceiveDate         string            `json:",omitempty"`
		CapturedAmount      uint32            `json:",omitempty"`
		CapturedDate        string            `json:",omitempty"`
		Currency            string            `json:",omitempty"`
		Country             string            `json:",omitempty"`
		ReturnCode          string            `json:",omitempty"`
		ReturnMessage       string            `json:",omitempty"`
		Status              uint32            `json:",omitempty"`
		Links               []*Links          `json:",omitempty"`
		ExtraDataCollection []interface{}     `json:",omitempty"`
		ExpirationDate      string            `json:",omitempty"`
		URL                 string            `json:",omitempty"`
		Number              string            `json:",omitempty"`
		BarCodeNumber       string            `json:",omitempty"`
		DigitableLine       string            `json:",omitempty"`
		Address             string            `json:",omitempty"`
	}

	Links struct {
		Method string `json:",omitempty"`
		Rel    string `json:",omitempty"`
		Href   string `json:",omitempty"`
	}

	RecurrentPayment struct {
		AuthorizeNow bool   `json:",omitempty"`
		EndDate      string `json:",omitempty"`
		Interval     string `json:",omitempty"`
	}

	CreditCard struct {
		CardNumber     string   `json:",omitempty"`
		CustomerName   string   `json:",omitempty"`
		Holder         string   `json:",omitempty"`
		ExpirationDate string   `json:",omitempty"`
		SecurityCode   string   `json:",omitempty"`
		SaveCard       bool     `json:",omitempty"`
		Brand          string   `json:",omitempty"`
		CardToken      string   `json:",omitempty"`
		Links          []*Links `json:"-"`
	}

	DebitCard struct {
		CardNumber                  string           `json:",omitempty"`
		CustomerName                string           `json:",omitempty"`
		Authenticate                bool             `json:",omitempty"`
		ReturnUrl                   string           `json:",omitempty"`
		IsCryptoCurrencyNegotiation bool             `json:",omitempty"`
		Holder                      string           `json:",omitempty"`
		ExpirationDate              string           `json:",omitempty"`
		RecurrentPayment            RecurrentPayment `json:",omitempty"`
		SecurityCode                string           `json:",omitempty"`
		CardOnFile                  CardOnFile       `json:",omitempty"`
		Brand                       string           `json:",omitempty"`
		CardToken                   string           `json:",omitempty"`
		Links                       []*Links         `json:",omitempty"`
	}

	CardOnFile struct {
		Usage  string `json:",omitempty"`
		Reason string `json:",omitempty"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
