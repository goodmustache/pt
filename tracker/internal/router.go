package internal

import (
	"net/http"

	"github.com/tedsuo/rata"
)

// Naming convention:
//
// Method + non-parameter parts of the path
//
// If the request returns a single entity by GUID, use the singular (for example
// /projects/{project_id} is GetProject).
//
// The const name should always be the const value + Request.
const (
	GetMeRequest       = "GetMe"
	GetProjectsRequest = "GetProjects"
)

// APIRoutes is a list of routes used by the rata library to construct request
// URLs.
var APIRoutes = rata.Routes{
	{Path: "/me", Method: http.MethodGet, Name: GetMeRequest},
	{Path: "/projects", Method: http.MethodGet, Name: GetProjectsRequest},
}
