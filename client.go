package strava

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Strava URL
const HostURL string = "https://www.strava.com/api/v3"

// Client -
type Client struct {
	HostURL      string
	HTTPClient   *http.Client
	ClientId     string
	ClientSecret string
}

// NewClient -
func NewClient(host, clientId, clientSecret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Strava URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If clientId or clientSecret not provided, return empty client
	if clientId == nil || clientSecret == nil {
		return &c, nil
	}

	c.ClientId = *clientId
	c.ClientSecret = *clientSecret

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	q := req.URL.Query()
	q.Set("client_id", c.ClientId)
	q.Set("client_secret", c.ClientSecret)
	req.URL.RawQuery = q.Encode()

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
