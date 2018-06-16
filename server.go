package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}
	http.HandleFunc("/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		//err = handlePost(w, r)
		//case: "PUT"
		//err = handlePut(w,r)
		//case:"DELETE"
		//err = handleDelete(w,r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "こんにちは。こんにちは。")
	return nil
}
