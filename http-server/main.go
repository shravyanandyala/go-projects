package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	connPort string = "8080"
)

type HttpHandler struct {
}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	buf, _ := ioutil.ReadAll(req.Body)
	fmt.Println(buf)
	res.Write(buf)
}

func main() {
	handler := HttpHandler{}
	http.ListenAndServe(":"+connPort, handler)
}
