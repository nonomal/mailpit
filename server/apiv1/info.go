package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/axllent/mailpit/internal/stats"
)

// AppInfo returns some basic details about the running app, and latest release.
func AppInfo(w http.ResponseWriter, _ *http.Request) {
	// swagger:route GET /api/v1/info application AppInformation
	//
	// # Get application information
	//
	// Returns basic runtime information, message totals and latest release version.
	//
	//	Produces:
	//	- application/json
	//
	//	Schemes: http, https
	//
	//	Responses:
	//		200: InfoResponse
	//		default: ErrorResponse

	info := stats.Load()

	bytes, _ := json.Marshal(info)

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(bytes)
}
