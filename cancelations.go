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
	if err != nil {
		return salePayment, err
	}
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
	if err != nil {
		return salePayment, err
	}
	return salePayment, nil
}

//CancelRecurrentPayment returns the payment/sale of cancel by PaymentId
//Endpoint PUT /1/RecurrentPayment/{RecurrentPaymentId}/Payment
func (c *Client) CancelRecurrentPayment(paymentId string) error {
	buf := bytes.NewBuffer([]byte(""))
	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s%s%s", c.Environment.APIUrl, "/1/RecurrentPayment/", paymentId, "/Deactivate"), buf)

	if err != nil {
		return err
	}

	err = c.Send(req, nil)
	if err != nil {
		return err
	}
	return nil
}
