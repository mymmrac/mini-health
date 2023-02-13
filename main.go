package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	quiteFlag := flag.Bool("q", false, "Quite output (only errors)")
	baseURLFlag := flag.String("e", "", "Environment variable name that contains base URL")
	statusCodeFlag := flag.Int("c", http.StatusBadRequest, "Smallest error HTTP status code")
	methodFlag := flag.String("m", http.MethodGet, "HTTP method used for request")
	dataFlag := flag.String("d", "", "Data passed as body of request")

	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n  %s [OPTIONS] URL\n\nOptions:\n", os.Args[0])
		flag.CommandLine.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		_, _ = fmt.Fprint(os.Stderr, "\nExpected exactly one argument (the health check URL)\n")
		os.Exit(1)
	}

	url := flag.Args()[0]
	if *baseURLFlag != "" {
		baseURL, ok := os.LookupEnv(*baseURLFlag)
		if !ok {
			_, _ = fmt.Fprintf(os.Stderr, "Environment varable %q not found\n", *baseURLFlag)
			os.Exit(1)
		}

		url = baseURL + url
	}

	req, err := http.NewRequest(*methodFlag, url, strings.NewReader(*dataFlag))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Create request failed: %s\n", err)
		os.Exit(1)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Request: %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode >= *statusCodeFlag {
		_, _ = fmt.Fprintf(os.Stderr, "Bad HTTP status code: %s\n", resp.Status)
		os.Exit(1)
	}

	if !*quiteFlag {
		fmt.Println(resp.Status)
	}
}
