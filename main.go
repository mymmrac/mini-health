package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprint(os.Stderr, "Expected exactly one argument (the health check URL)\n")
		os.Exit(1)
	}

	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Request: %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		_, _ = fmt.Fprintf(os.Stderr, "Bad HTTP status code: %s\n", resp.Status)
		os.Exit(1)
	}
}
