package strava

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// GetAllSubscriptions - Returns all subscriptions
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

// GetSubscription - Returns a subscription
func (c *Client) GetSubscription(subscriptionID int) (*Subscription, error) {
	subscriptions, err := c.GetAllSubscriptions()
	if err != nil {
		return nil, err
	}

	for _, subscription := range *subscriptions {
		if subscription.ID == subscriptionID {
			return &subscription, nil
		}
	}

	return nil, errors.New("Subscription not found")
}

// CreateSubscription - Create a new subscription
func (c *Client) CreateSubscription(subscriptionItem SubscriptionItem) (*Subscription, error) {
	url, err := url.Parse(c.HostURL)
	if err != nil {
		return nil, err
	}

	url = url.JoinPath("push_subscriptions")

	q := url.Query()
	q.Set("callback_url", subscriptionItem.CallbackURL)
	q.Set("verify_token", subscriptionItem.VerifyToken)
	url.RawQuery = q.Encode()

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
