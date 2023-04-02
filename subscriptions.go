package strava

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// GetAllSubscriptions - Returns all subscription details
func (c *Client) GetAllSubscriptions() (*[]Subscription, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/push_subscriptions", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	subscriptions := []Subscription{}
	err = json.Unmarshal(body, &subscriptions)
	if err != nil {
		return nil, err
	}

	return &subscriptions, nil
}

// CreateSubscription - Create a new subscription
func (c *Client) CreateSubscription(subscriptionItem SubscriptionItem) (*Subscription, error) {
	url := url.URL{
		Host: c.HostURL,
		Path: "/push_subscriptions",
	}

	q := url.Query()
	q.Add("callback_url", subscriptionItem.CallbackURL)
	q.Add("verify_token", subscriptionItem.VerifyToken)

	req, err := http.NewRequest("POST", url.String(), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	subscription := Subscription{}
	err = json.Unmarshal(body, &subscription)
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

// DeleteSubscription - Deletes a subscription
func (c *Client) DeleteSubscription(subscriptionID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/subscriptions/%s", c.HostURL, subscriptionID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted subscription" {
		return errors.New(string(body))
	}

	return nil
}
