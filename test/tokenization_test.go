package test

import (
	"github.com/vasconcelosvcd/go-cielo"
	"testing"
)

func TestTokenize(t *testing.T) {
	c, err := getNewClient(true, t)
	if err != nil {
		t.Error(err.Error())
	}
	if c == nil {
		t.Fatal("Client must not be nil")
	}

	cc := cielo.CreditCard{
		CardNumber:     "5247712516640978",
		CustomerName:   "Tester Name",
		Holder:         "Tester Holder",
		ExpirationDate: "11/2026",
		SaveCard:       true,
		Brand:          "Master",
	}

	token, err := c.CreateTokenizeCard(&cc)
	if err != nil {
		t.Error(err.Error())
	}
	if token == nil {
		t.Fatal("Token must not be nil")
	}
	if len(token.CardToken) <= 0 {
		t.Error("CC Token was not generated")
	}
}

func TestGetTokenized(t *testing.T) {
	c, err := getNewClient(true, t)
	if err != nil {
		t.Error(err.Error())
	}
	if c == nil {
		t.Fatal("Client must not be nil")
	}

	cc := cielo.CreditCard{
		CardNumber:     "5247712516640978",
		CustomerName:   "Tester Name",
		Holder:         "Tester Holder",
		ExpirationDate: "11/2026",
		SaveCard:       true,
		Brand:          "Master",
	}

	token, err := c.CreateTokenizeCard(&cc)
	if err != nil {
		t.Error(err.Error())
	}
	if token == nil {
		t.Fatal("Token must not be nil")
	}
	if len(token.CardToken) <= 0 {
		t.Error("CC Token was not generated")
	}

	tokenizedCard, err := c.GetTokenizeCard(token.CardToken)
	if err != nil {
		t.Error(err.Error())
	}
	if tokenizedCard == nil {
		t.Fatal("Tokenized card must not be nil")
	}
	if len(tokenizedCard.CardNumber) <= 0 {
		t.Error("CC Token was not generated")
	}
}
