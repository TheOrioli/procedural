package endpoints

import "net/http"

// Route contains endpoint information
type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}
