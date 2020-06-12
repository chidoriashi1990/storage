/*
 * Storage Scales
 *
 * 💪💪 It is an application that calculates the size of files under a directory and returns them in JSON format.
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package scales

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route model
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type
type Routes []Route

// NewRouter func
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index func
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"GetWeight",
		strings.ToUpper("Post"),
		"/v1/weight",
		GetWeight,
	},
}
