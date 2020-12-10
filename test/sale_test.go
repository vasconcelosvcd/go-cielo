package test

import (
	"github.com/vasconcelosvcd/go-cielo"
	"testing"
)

func Test_SaleCreditCard(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		c, err := getNewClient(true, t)

		if err != nil {
			t.Error(err.Error())
		}
		if c == nil {
			t.Fatal("Client must not be nil")
		}

		sale := cielo.Sale{
			MerchantOrderID: "12387197ads89d7a",
			Customer: &cielo.Customer{
				Name: "Customer Test",
			},
			Payment: &cielo.Payment{
				SoftDescriptor: "Simple Sale",
				Type:           "CreditCard",
				Installments:   1,
				Amount:         1300,
				CreditCard: &cielo.CreditCard{
					CardNumber:     "5247712516640978",
					CustomerName:   "Tester Name",
					Holder:         "Tester Holder",
					ExpirationDate: "11/2026",
					SaveCard:       false,
					Brand:          "Master",
				},
			},
		}

		saleReturn, err := c.Authorization(&sale)
		if err != nil {
			t.Fatal(err.Error())
		}
		if saleReturn == nil {
			t.Fatal(" Sale must not be nil")
		}
		if len(saleReturn.Payment.PaymentID) <= 0 {
			t.Fatal("payment was not generated")
		}
	})

}
