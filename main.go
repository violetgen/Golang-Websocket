package main

import (
	"net/http"
)

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//upgrader switch the http to websocket
func main() {
	// router := &Router{}
	router := NewRouter()

	router.Handle("channel add", addChannel)
	http.Handle("/", router)
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
