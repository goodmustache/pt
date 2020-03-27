package tracker_test

import (
	"net/http"

	. "github.com/goodmustache/pt/tracker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Project", func() {
	var client *Client

	BeforeEach(func() {
		client, _ = NewTestClient()
	})

	Describe("Projects", func() {
		var (
			projects   []Project
			executeErr error
		)

		JustBeforeEach(func() {
			projects, executeErr = client.Projects()
		})

		Context("when a successful response is returned", func() {
			BeforeEach(func() {
				response := `[
					{
						"id": 122345,
						"kind": "project",
						"name": "Test project 1",
						"version": 2867,
						"iteration_length": 1,
						"week_start_day": "Monday",
						"point_scale": "0,1,2,3,5,8",
						"point_scale_is_custom": false,
						"bugs_and_chores_are_estimatable": false,
						"automatic_planning": true,
						"enable_tasks": true,
						"time_zone": {
							"kind": "time_zone",
							"olson_name": "America/Los_Angeles",
							"offset": "-07:00"
						},
						"velocity_averaged_over": 3,
						"number_of_done_iterations_to_show": 4,
						"has_google_domain": false,
						"description": "This is the description for this project",
						"enable_incoming_emails": true,
						"initial_velocity": 10,
						"public": false,
						"atom_enabled": false,
						"project_type": "shared",
						"start_time": "2019-07-22T07:00:00Z",
						"created_at": "2019-07-22T16:53:10Z",
						"updated_at": "2020-03-10T18:56:43Z",
						"account_id": 5432101,
						"current_iteration_number": 35,
						"enable_following": true
					},
					{
						"id": 152649,
						"kind": "project",
						"name": "Test project 2",
						"version": 2867,
						"iteration_length": 1,
						"week_start_day": "Monday",
						"point_scale": "0,1,2,3,5,8",
						"point_scale_is_custom": false,
						"bugs_and_chores_are_estimatable": false,
						"automatic_planning": true,
						"enable_tasks": true,
						"time_zone": {
							"kind": "time_zone",
							"olson_name": "America/Los_Angeles",
							"offset": "-07:00"
						},
						"velocity_averaged_over": 3,
						"number_of_done_iterations_to_show": 4,
						"has_google_domain": false,
						"description": "different project description",
						"enable_incoming_emails": true,
						"initial_velocity": 10,
						"public": true,
						"atom_enabled": false,
						"project_type": "shared",
						"start_time": "2019-07-22T07:00:00Z",
						"created_at": "2019-07-22T16:53:10Z",
						"updated_at": "2020-03-10T18:56:43Z",
						"account_id": 54321010,
						"current_iteration_number": 35,
						"enable_following": true
					}
				]`

				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/projects"),
						RespondWith(http.StatusAccepted, response),
					),
				)
			})

			It("returns list of project", func() {
				Expect(executeErr).ToNot(HaveOccurred())
				Expect(projects).To(ConsistOf(
					Project{
						Description: "This is the description for this project",
						ID:          122345,
						Name:        "Test project 1",
						Public:      false,
					},
					Project{
						Description: "different project description",
						ID:          152649,
						Name:        "Test project 2",
						Public:      true,
					},
				))
			})
		})
	})
})
