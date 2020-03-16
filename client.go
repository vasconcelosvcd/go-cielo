package cielo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func NewClient(merchantId string, merchantKey string, environment Environment) (*Client, error) {
	if merchantId == "" || merchantKey == "" || environment.APIUrl == "" || environment.APIQueryURL == "" {
		return nil, errors.New("MerchantID, MerchantKey and environment are required to create a Client")
	}

	return &Client{
		Client:      &http.Client{},
		Environment: environment,
		MerchantId:  merchantId,
		MerchantKey: merchantKey,
	}, nil
}

// Send makes a request to the API, the response body will be
// unmarshaled into v, or if v is an io.Writer, the response will
// be written to it without decoding
func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	//// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")
	//Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "CieloEcommerce/3.0 GoCielo")
	req.Header.Add("MerchantId", c.MerchantId)
	req.Header.Add("MerchantKey", c.MerchantKey)
	req.Header.Add("RequestId", uuid.NewV5(uuid.NamespaceX500, "go-cielo").String())
	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			_ = json.Unmarshal(data, errResp)
		}

		return errResp
	}
	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, _ = io.Copy(w, resp.Body)
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// log will dump request and response to the log file
func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		_, _ = c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}

// NewRequest constructs a request
// Convert payload to a JSON
func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	return http.NewRequest(method, url, buf)
}
