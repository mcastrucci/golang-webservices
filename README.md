Go Web service Boilerplate
==========================

This is a little util framework that can be used to create web services / REST APIs faster using Go
the idea is to import the lib and use the "StartServer" function. It accepts string and int parameters to define the port.

before calling StartServer endpoints should be added using func AddEndpoint.
It takes as parameter an url and a handler. Handler should be child of 
type FunctionHandler = func(w http.ResponseWriter, r *http.Request)
