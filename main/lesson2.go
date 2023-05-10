package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var version string

func main() {
	os.Setenv("VERSION", "1.1.0")
	version = os.Getenv("VERSION")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", healthzHandler)
	http.ListenAndServe(":80", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		for i := 0; i < len(v); i ++ {
			res.Header().Add(k, v[0])
		}
	}
	res.Header().Add("VERSION", version)

	fmt.Fprintf(res, "hello world")

	hostPort := strings.Split(req.Host, ":")
	fmt.Printf("ip:{%s}, http code: %d \n", hostPort[0], 200)
}

func healthzHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "200")

	hostPort := strings.Split(req.Host, ":")
	fmt.Printf("ip:{%s}, http code: %d \n", hostPort[0], 200)
}