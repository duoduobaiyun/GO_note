package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index",ind)
	err:=http.ListenAndServe("localhost:8090",nil)
	fmt.Println(err)
}
func ind(w http.ResponseWriter,r *http.Request) {
	w.Write([]byte("你好"))
}
