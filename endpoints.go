/*
	This file is destinated to put the endpoints that the web service or REST Api will use
	They should be of type "Handler" and be defined in the endpoints.json in order to use
	if endpoints.json does not exist, create one with the following sintax
	{
		"endpoints": [
			{
				"url": "/",
				"handler": "HomeHandler"
			}
		]
	}
*/
package webservices

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
