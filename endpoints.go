/*
	This is a template file that show how a handler should be
	in order to be accepted as endpoint
*/
package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage! The app works")
}
