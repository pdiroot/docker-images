package main

import (
	"fmt"
	"net/http"
	"time"
)

func report(w http.ResponseWriter, req *http.Request) {
	time.Sleep(75 * time.Second)
	fmt.Fprintf(w, "Report completed\n")
}

func main() {
	http.HandleFunc("/report", report)
	http.ListenAndServe(":80", nil)
}
