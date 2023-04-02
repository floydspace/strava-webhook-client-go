package strava

// Subscription -
type Subscription struct {
	ID            int    `json:"id,omitempty"`
	ResourceState int    `json:"resource_state,omitempty"`
	ApplicationID int    `json:"application_id,omitempty"`
	CallbackURL   string `json:"callback_url,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

// SubscriptionItem -
type SubscriptionItem struct {
	CallbackURL string `json:"callback_url"`
	VerifyToken string `json:"verify_token"`
}
