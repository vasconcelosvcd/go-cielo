package cielo

import (
	"bytes"
	"fmt"
)

//CancelByPaymentId returns the payment/sale cancel by paymentId
//Endpoint POST /1/sales/{payment_id}/void
func (c *Client) CancelByPaymentId(paymentId string) (*Payment, error) {
	buf := bytes.NewBuffer([]byte(""))
	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s%s", c.Environment.APIUrl, "/1/sales/", paymentId), buf)

	salePayment := &Payment{}

	if err != nil {
		return salePayment, err
	}

	err = c.Send(req, salePayment)
	return salePayment, nil
}

//CancelByPaymentId returns the payment/sale cancel by merchantOrderId
//Endpoint POST /1/sales/{merchantOrder_id}/void
func (c *Client) CancelByMerchantOrderId(mOrderId string) (*Payment, error) {
	buf := bytes.NewBuffer([]byte(""))
	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s%s", c.Environment.APIUrl, "/1/sales/", mOrderId), buf)

	salePayment := &Payment{}

	if err != nil {
		return salePayment, err
	}

	err = c.Send(req, salePayment)
	return salePayment, nil
}
