package goadafruit

import (
	"fmt"
	"time"
)

type User struct {
	User struct {
		ID                 int         `json:"id"`
		Name               string      `json:"name"`
		Color              string      `json:"color"`
		Username           string      `json:"username"`
		Role               string      `json:"role"`
		DefaultGroupID     int         `json:"default_group_id"`
		DefaultDashboardID interface{} `json:"default_dashboard_id"`
		TimeZone           string      `json:"time_zone"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		Subscription       struct {
			Status string `json:"status"`
			Limits struct {
				Feeds      int `json:"feeds"`
				Dashboards int `json:"dashboards"`
				Groups     int `json:"groups"`
				DataTTL    int `json:"data_ttl"`
				DataRate   int `json:"data_rate"`
			} `json:"limits"`
			Plan struct {
				Name       string `json:"name"`
				Price      int    `json:"price"`
				Interval   string `json:"interval"`
				StripeID   string `json:"stripe_id"`
				BaseLimits struct {
					Feeds      int `json:"feeds"`
					Dashboards int `json:"dashboards"`
					Groups     int `json:"groups"`
					DataTTL    int `json:"data_ttl"`
					DataRate   int `json:"data_rate"`
				} `json:"base_limits"`
				Free bool `json:"free"`
				Paid bool `json:"paid"`
			} `json:"plan"`
		} `json:"subscription"`
		BetaFlags interface{} `json:"beta_flags"`
	} `json:"user"`
	Profile struct {
		Name              string        `json:"name"`
		Username          string        `json:"username"`
		TimeZone          string        `json:"time_zone"`
		CreatedAt         time.Time     `json:"created_at"`
		ConnectedAccounts []interface{} `json:"connected_accounts"`
		Subscription      struct {
			ID   interface{} `json:"id"`
			Plan struct {
				Name       string  `json:"name"`
				Price      float64 `json:"price"`
				Interval   string  `json:"interval"`
				StripeID   string  `json:"stripe_id"`
				BaseLimits struct {
					Feeds      int `json:"feeds"`
					Dashboards int `json:"dashboards"`
					Groups     int `json:"groups"`
					DataTTL    int `json:"data_ttl"`
					DataRate   int `json:"data_rate"`
				} `json:"base_limits"`
				Free bool `json:"free"`
				Paid bool `json:"paid"`
			} `json:"plan"`
			CreatedAt interface{} `json:"created_at"`
			Status    string      `json:"status"`
			Limits    struct {
				Feeds      int `json:"feeds"`
				Dashboards int `json:"dashboards"`
				Groups     int `json:"groups"`
				DataTTL    int `json:"data_ttl"`
				DataRate   int `json:"data_rate"`
			} `json:"limits"`
			Price    int           `json:"price"`
			Upgrades []interface{} `json:"upgrades"`
		} `json:"subscription"`
		HasBillingHistory bool        `json:"has_billing_history"`
		HasPaymentSource  bool        `json:"has_payment_source"`
		AccountBalance    int         `json:"account_balance"`
		BetaDiscount      interface{} `json:"beta_discount"`
		ActionItems       struct {
			PendingShares []interface{} `json:"pending_shares"`
		} `json:"action_items"`
		PromotionDiscounts []interface{} `json:"promotion_discounts"`
		Coupons            []interface{} `json:"coupons"`
	} `json:"profile"`
	Sidebar struct {
		FeedCount      int `json:"feed_count"`
		GroupCount     int `json:"group_count"`
		DashboardCount int `json:"dashboard_count"`
		ActiveDataRate int `json:"active_data_rate"`
	} `json:"sidebar"`
	Navigation struct {
		Feeds struct {
			Records []struct {
				ID        int       `json:"id"`
				Name      string    `json:"name"`
				Key       string    `json:"key"`
				UpdatedAt time.Time `json:"updated_at"`
				LastValue string    `json:"last_value"`
			} `json:"records"`
			Count int `json:"count"`
		} `json:"feeds"`
		Dashboards struct {
			Records []struct {
				ID        int       `json:"id"`
				Name      string    `json:"name"`
				UpdatedAt time.Time `json:"updated_at"`
				Key       string    `json:"key"`
			} `json:"records"`
			Count int `json:"count"`
		} `json:"dashboards"`
		Devices struct {
			Records []interface{} `json:"records"`
		} `json:"devices"`
		Triggers struct {
			Records []interface{} `json:"records"`
		} `json:"triggers"`
	} `json:"navigation"`
	Throttle struct {
		DataRateLimit              int         `json:"data_rate_limit"`
		ActiveDataRate             int         `json:"active_data_rate"`
		AuthenticationRate         int         `json:"authentication_rate"`
		SubscribeAuthorizationRate int         `json:"subscribe_authorization_rate"`
		PublishAuthorizationRate   int         `json:"publish_authorization_rate"`
		HourlyBanRate              int         `json:"hourly_ban_rate"`
		MqttBanErrorMessage        interface{} `json:"mqtt_ban_error_message"`
	} `json:"throttle"`
	SystemMessages []interface{} `json:"system_messages"`
}

type UserDetails struct {
	DataRateLimit              int         `json:"data_rate_limit"`
	ActiveDataRate             int         `json:"active_data_rate"`
	AuthenticationRate         int         `json:"authentication_rate"`
	SubscribeAuthorizationRate int         `json:"subscribe_authorization_rate"`
	PublishAuthorizationRate   int         `json:"publish_authorization_rate"`
	HourlyBanRate              int         `json:"hourly_ban_rate"`
	MqttBanErrorMessage        interface{} `json:"mqtt_ban_error_message"`
}

type UserService struct {
	client *Client
}

func (s *UserService) GetUserInformation() (*User, *Response, error) {
	path := "/api/v2/user"

	req, rerr := s.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var user User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}

func (u *UserService) GetDetailedUserInformation() (*UserDetails, *Response, error) {
	path := fmt.Sprintf("/api/v2/%v/throttle", u.client.Username)

	req, rerr := u.client.NewRequest("GET", path, nil)
	if rerr != nil {
		return nil, nil, rerr
	}

	var user UserDetails
	resp, err := u.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}
