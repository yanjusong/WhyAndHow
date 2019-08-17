package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"
)

func getHash(bytes []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(bytes))
}

// GenerateETag generates a new ETag string by the specified content
func GenerateETag(bytes []byte, weak bool) string {
	tag := fmt.Sprintf("\"%d-%s\"", len(bytes), getHash(bytes))
	if weak {
		tag = "W/" + tag
	}

	return tag
}

func commonHandler(w http.ResponseWriter, r *http.Request) {
	// 获取分钟级别的时间字符串
	timeStr := time.Now().Format("2006-01-02 15:04")

	reqETag := r.Header.Get("If-None-Match")
	curETag := GenerateETag([]byte(timeStr), false)

	if reqETag != "" {
		// ETag match, response 304 indicating not modified
		if curETag == reqETag {
			fmt.Printf("response 304\n")
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	w.Header().Set("ETag", curETag)
	w.Write([]byte(timeStr))
}

func main() {
	http.HandleFunc("/time", commonHandler)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("starting server failed: %s\n", err.Error())
	}
}
