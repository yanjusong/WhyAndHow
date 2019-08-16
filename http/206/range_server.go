package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const data  = "0123456789"

func respBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("lack of \"Range\" header, it must contains a header like \"Range: bytes=0-9\", " +
		"the left is greater than or equal to 0, and the right is less than or equal to 9, " +
		"and the left must less than or equal to the right\n"))
}

func rangeHandler(w http.ResponseWriter, r *http.Request) {
	rangeCommand := r.Header.Get("Range")
	if rangeCommand == "" {
		respBadRequest(w)
		return
	}

	s := strings.Split(rangeCommand, "=")
	if len(s) != 2 {
		respBadRequest(w)
		return
	}

	rangeNums := strings.Split(s[1], "-")
	if len(rangeNums) != 2 {
		respBadRequest(w)
		return
	}

	left, err := strconv.Atoi(strings.Trim(rangeNums[0], " "))
	if err != nil {
		respBadRequest(w)
		return
	}

	right, err := strconv.Atoi(strings.Trim(rangeNums[1], " "))
	if err != nil {
		respBadRequest(w)
		return
	}

	if left < 0 || left > 9 || right < 0 || right > 9 || left > right {
		respBadRequest(w)
		return
	}

	respData := data[left:right+1]

	w.Header().Set("Content-Range", fmt.Sprint("bytes %d-%d/%d", left, right, len(data)))
	w.WriteHeader(http.StatusPartialContent)
	w.Write([]byte(respData))
}

func main() {
	http.HandleFunc("/range", rangeHandler)
	http.ListenAndServe(":9090", nil)
}
