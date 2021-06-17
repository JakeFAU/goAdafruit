package goadafruit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInformation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/user",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
			{
				"user": {
				  "id": 509745,
				  "name": "Jacob Bourne",
				  "color": "#1433e3",
				  "username": "JakeFau",
				  "role": "user",
				  "default_group_id": 435375,
				  "default_dashboard_id": null,
				  "time_zone": "America/New_York",
				  "created_at": "2021-05-14T01:59:27Z",
				  "updated_at": "2021-06-17T18:07:14Z",
				  "subscription": {
					"status": "active",
					"limits": {
					  "feeds": 10,
					  "dashboards": 5,
					  "groups": 5,
					  "data_ttl": 2592000,
					  "data_rate": 30
					},
					"plan": {
					  "name": "Free",
					  "price": 0,
					  "interval": "month",
					  "stripe_id": "io-free",
					  "base_limits": {
						"feeds": 10,
						"dashboards": 5,
						"groups": 5,
						"data_ttl": 2592000,
						"data_rate": 30
					  },
					  "free": true,
					  "paid": false
					}
				  },
				  "beta_flags": null
				},
				"profile": {
				  "name": "Jacob Bourne",
				  "username": "JakeFau",
				  "time_zone": "America/New_York",
				  "created_at": "2021-05-14T01:59:27Z",
				  "connected_accounts": [],
				  "subscription": {
					"id": null,
					"plan": {
					  "name": "Free",
					  "price": 0,
					  "interval": "month",
					  "stripe_id": "io-free",
					  "base_limits": {
						"feeds": 10,
						"dashboards": 5,
						"groups": 5,
						"data_ttl": 2592000,
						"data_rate": 30
					  },
					  "free": true,
					  "paid": false
					},
					"created_at": null,
					"status": "free",
					"limits": {
					  "feeds": 10,
					  "dashboards": 5,
					  "groups": 5,
					  "data_ttl": 2592000,
					  "data_rate": 30
					},
					"price": 0,
					"upgrades": []
				  },
				  "has_billing_history": false,
				  "has_payment_source": false,
				  "account_balance": 0,
				  "beta_discount": null,
				  "action_items": {
					"pending_shares": []
				  },
				  "promotion_discounts": [],
				  "coupons": []
				},
				"sidebar": {
				  "feed_count": 5,
				  "group_count": 3,
				  "dashboard_count": 1,
				  "active_data_rate": 0
				},
				"navigation": {
				  "feeds": {
					"records": [
					  {
						"id": 1638924,
						"name": "Weather / Humidity",
						"key": "weather.humidity",
						"updated_at": "2021-06-16T16:04:09Z",
						"last_value": "5.10e+01"
					  },
					  {
						"id": 1640907,
						"name": "Outdoor / Temperature",
						"key": "outdoor.temperature",
						"updated_at": "2021-06-16T16:04:09Z",
						"last_value": "28.9"
					  },
					  {
						"id": 1640909,
						"name": "Outdoor / Humidity",
						"key": "outdoor.humidity",
						"updated_at": "2021-06-16T16:04:09Z",
						"last_value": "75.0"
					  },
					  {
						"id": 1638911,
						"name": "Weather / Temperature",
						"key": "weather.temperature",
						"updated_at": "2021-06-16T16:04:09Z",
						"last_value": "7.56e+01"
					  },
					  {
						"id": 1638925,
						"name": "Weather / Pressure",
						"key": "weather.pressure",
						"updated_at": "2021-06-16T16:04:09Z",
						"last_value": "3.00e+01"
					  }
					],
					"count": 5
				  },
				  "dashboards": {
					"records": [
					  {
						"id": 445654,
						"name": "Jake's Apt",
						"updated_at": "2021-06-17T09:33:07Z",
						"key": "jakes-apt"
					  }
					],
					"count": 1
				  },
				  "devices": {
					"records": []
				  },
				  "triggers": {
					"records": []
				  }
				},
				"throttle": {
				  "data_rate_limit": 30,
				  "active_data_rate": 0,
				  "authentication_rate": 0,
				  "subscribe_authorization_rate": 0,
				  "publish_authorization_rate": 0,
				  "hourly_ban_rate": 0,
				  "mqtt_ban_error_message": null
				},
				"system_messages": []
			  }`)
		})

	assert := assert.New(t)

	user, response, err := client.User.GetUserInformation()

	assert.Nil(err)
	assert.NotNil(user)
	assert.NotNil(response)

	assert.Equal("Jacob Bourne", user.User.Name)

}

func TestDetailedUserInformation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2/test-user/throttle",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, `
		{
			"data_rate_limit": 60,
			"active_data_rate": 20
		}`)
		})

	assert := assert.New(t)

	user, response, err := client.User.GetDetailedUserInformation()

	assert.Nil(err)
	assert.NotNil(user)
	assert.NotNil(response)

	assert.Equal(60, user.DataRateLimit)

}
