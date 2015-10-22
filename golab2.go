package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MyJsonName struct {
	Name string `json:"name"`
}

type MyJsonResponse struct {
	Greeting string
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func hellopost(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	var myjson3 MyJsonName
	s3 := json.NewDecoder(req.Body)
	err := s3.Decode(&myjson3)

	var myJsonresp2 MyJsonResponse
	myJsonresp2.Greeting = "Hello, " + myjson3.Name + "!"
	b2, err := json.Marshal(myJsonresp2)
	if err != nil {
	}
	fmt.Fprintf(rw, string(b2))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", hellopost)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
