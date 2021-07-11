/*
	This File Starts the server in different wayss
*/

package webservices

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

var DEFAULT_PORT = 8080

/*
	It starts the Server. It accepts as parameter an integer for port number
	or Args that would contain port number.
	If no port is provided. The server will start in default 8080 port
*/
func StartServer(parameters ...interface{}) (string, error) {
	var port int
	var error error
	var fail = true

	if len(parameters) > 0 {
		/*
			for the moment, we only accept one paramenter, so we will take the first arg from the
			params
		*/
		args := parameters[0]

		argType := reflect.TypeOf(args).Kind()

		switch argType {
		case reflect.String:
			/*
				When main args are sent as parameters, the first arg is the path
				of the application. We should give it another loop
			*/
			parsedPort, err := strconv.Atoi(args.(string))
			if err != nil {
				error = errors.New("Invalid Argument, port should be numeric and less than 65535, server not started")
			} else {
				port = parsedPort
				fail = false
			}
		case reflect.Int:
			port = args.(int)
			fail = false
		default:
			error = errors.New("invalid parameters, only port number (int or string)")
		}

		/*
			After checking the args, we start the server or return error
			if loop failed to get a correct arg
		*/
		if fail {
			return "", error
		} else {
			return startServer(port)
		}

	} else {
		// No params, starting the server with default port
		return startServer(DEFAULT_PORT)
	}
}

/*
	Private Function that will start the server
	with the port passed by parameter.
	It returns a debug String for success and an error in case of fail
*/
func startServer(port int) (string, error) {

	var error error

	log.Print(fmt.Sprintf("Starting server at port %d \n", port))

	for _, v := range allowedEndpoints {
		log.Printf("Listening for endpoint %s \n", v)

		http.HandleFunc(v, endPointHandler)
	}
	// http.HandleFunc("/", endPointHandler)
	portListening := fmt.Sprintf(":%d", port)
	error = http.ListenAndServe(portListening, nil)
	return "Server Started succesfully", error
}
