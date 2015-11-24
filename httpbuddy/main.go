package main

import (
	"flag"
	"fmt"
)

var port int

func main() {
	flag.IntVar(&port, "port", 3000, "The port to run the server on")
	flag.Parse()

	fmt.Print("Serving files on localhost:%v\n", port)
}

func ServeStatic(port int) error {
	host := fmt.Sprint("localhost: %v", port)
	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}
