package cielo

import (
	"fmt"
)

//Sale returns the payment/sale authorization
//Endpoint POST /1/sales
func (c *Client) Authorization(sale *Sale) (*Sale, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Environment.APIUrl, "/1/sales/"), sale)

	salePayed := &Sale{}

	if err != nil {
		return salePayed, err
	}

	err = c.Send(req, salePayed)
	return salePayed, nil
}
