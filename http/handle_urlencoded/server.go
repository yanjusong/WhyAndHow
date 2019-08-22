package main

import (
	"fmt"
	"net/http"
)

func urlencodedHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Printf("name=%s\n", name)

	age := r.FormValue("age")
	fmt.Printf("age=%s\n", age)

	w.Write([]byte("you posted ok"))
}

func main() {
	http.HandleFunc("/post_urlencoded", urlencodedHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("starting server failed: %s\n", err.Error())
	}
}
