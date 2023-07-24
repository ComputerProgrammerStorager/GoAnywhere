package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer attempts to contact the server of a URL
// It tries for one minute using exponential back-off
// It reports an errof is all attempts fail

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying....", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to responsd after %s", url, timeout)
}

func main() {
	url := os.Args[1]
	log.SetPrefix("Wait:")
	log.SetFlags(0)
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Site %s is up and running\n", url)
}
