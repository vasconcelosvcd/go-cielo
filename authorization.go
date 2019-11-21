package cielo

import (
	"fmt"
)

//Sale returns the payment/sale authorization
//Endpoint POST /1/sales
func (c *Client) Authorization(payment *Payment) (*Payment, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Environment.APIUrl, "/1/sales/"), payment)

	salePayment := &Payment{}

	if err != nil {
		return salePayment, err
	}

	err = c.Send(req, salePayment)
	return salePayment, nil
}
