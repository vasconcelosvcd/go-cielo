package cielo

import (
	"bytes"
	"fmt"
)

//CreateTokenizeCard returns the tokenized card
//Endpoint POST /1/card
func (c *Client) CreateTokenizeCard(creditCard *CreditCard) (*CreditCard, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Environment.APIUrl, "/1/card/"), creditCard)

	cCard := &CreditCard{}

	if err != nil {
		return cCard, err
	}

	err = c.Send(req, creditCard)
	if err != nil {
		return cCard, err
	}
	return creditCard, err
}

//CreateTokenizeCard returns some card information
//Endpoint GET /1/card/{TOKEN}
func (c *Client) GetTokenizeCard(cardToken string) (*CreditCard, error) {
	buf := bytes.NewBuffer([]byte(""))
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.Environment.APIQueryURL, "/1/card/", cardToken), buf)

	cCard := &CreditCard{}

	if err != nil {
		return cCard, err
	}

	err = c.Send(req, cCard)
	return cCard, err
}
