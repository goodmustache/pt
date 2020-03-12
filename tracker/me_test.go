package tracker_test

import (
	"net/http"

	. "github.com/goodmustache/pt/tracker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Me", func() {
	var client *Client

	BeforeEach(func() {
		client, _ = NewTestClient()
	})

	Describe("Me", func() {
		var (
			me         Me
			executeErr error
		)

		JustBeforeEach(func() {
			me, executeErr = client.Me()
		})

		Context("when a successful response is returned", func() {
			BeforeEach(func() {
				response := `{
					"kind": "me",
					"id": 1010101,
					"name": "Anand Gaitonde",
					"initials": "❍ᴥ❍ʋ",
					"username": "XenoPhex",
					"time_zone": {
						"kind": "time_zone",
						"olson_name": "America/Chicago",
						"offset": "-06:00"
					},
					"api_token": "a013md020v9m2jc9ejc92m2n2s0qj200",
					"has_google_identity": true,
					"accounts": [
						{
							"kind": "account_summary",
							"id": 123456,
							"name": "gerrrate",
							"status": "active",
							"plan": "Free"
						}
					],
					"projects": [
						{
							"kind": "membership_summary",
							"id": 9876543,
							"project_id": 1000001,
							"project_name": "pt",
							"project_color": "b800bb",
							"favorite": false,
							"role": "owner",
							"last_viewed_at": "2020-01-03T02:05:50Z"
						}
					],
					"email": "test@example.com",
					"receives_in_app_notifications": true,
					"created_at": "2016-05-16T06:55:16Z",
					"updated_at": "2020-03-06T21:53:07Z"
				}`

				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/me"),
						RespondWith(http.StatusAccepted, response),
					),
				)
			})

			It("returns the deployment guid of the most recent deployment", func() {
				Expect(executeErr).ToNot(HaveOccurred())
				Expect(me).To(Equal(
					Me{
						APIToken: "a013md020v9m2jc9ejc92m2n2s0qj200",
						Email:    "test@example.com",
						ID:       1010101,
						Name:     "Anand Gaitonde",
						Username: "XenoPhex",
					},
				))
			})
		})
	})
})
