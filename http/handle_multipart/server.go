package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func multipartHandler(w http.ResponseWriter, r *http.Request) {
	for key, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			fmt.Println(key, value)
		}
	}

	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Printf("error:%+v\n", err)
	}

	for {
		p, err := reader.NextPart()

		if err == io.EOF {
			break
		}

		fileName := p.FileName()
		formName := p.FormName()
		formData, err := ioutil.ReadAll(p)
		if err != nil {
			fmt.Printf("error:%+v\n", err)
		}

		fmt.Printf("fileName=%s, formName=%s\n", fileName, formName)
		fmt.Printf("n=%d\n", len(formData))
	}
	w.Write([]byte("you posted ok"))
}

func main() {
	http.HandleFunc("/post_multipart", multipartHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("starting server failed: %s\n", err.Error())
	}
}
