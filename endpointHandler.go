package main

import (
	"net/http"
)

type FunctionHandler = func(w http.ResponseWriter, r *http.Request)

/*
	a map of url / functions that tha app will use
*/
var endpointsHandlers map[string]FunctionHandler = make(map[string]FunctionHandler)

/*
	An array of the Endpoints in the application
*/
var allowedEndpoints = make([]string, 0)

/*
	adds an endpoint to the list of allowed endpoints to be used in validations
	and security.
*/
func AddEndpoint(url string, handler FunctionHandler) {
	allowedEndpoints = append(allowedEndpoints, url)
	endpointsHandlers[url] = handler

}

/*
	An endpoint handler with some security implemented on it
	It will check form an array of endpoints if the endpoint and method is available or not
	and if we have a handler for it
*/
func endPointHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	/*
		We check if the url requested is available in our application
		if it does not exist or user is not allowed, we reject with statusNotFound
	*/
	var validUrl = ContainsUrl(allowedEndpoints, url)

	if validUrl {
		/*
			We also need to check if we have a valid handler for it
		*/
		handler, exist := endpointsHandlers[url]
		if exist {
			handler(w, r) // if it exists, we just call it
		} else {
			http.Error(w, "invalid direction.", http.StatusNotFound)
		}
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}
