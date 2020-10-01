package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("开心就好"))
	})

	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal(err)
		return
	}
}
