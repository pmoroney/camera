package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type logHandler struct{}

func (lh *logHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	log.Printf("Req: %s %s%s, %v, %s", req.Method, req.Host, req.RequestURI, req.Header, body)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.ListenAndServe(":8023", &logHandler{})
}
