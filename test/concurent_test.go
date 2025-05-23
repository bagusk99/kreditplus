package test

import (
	"net/http"
	"testing"
	"time"
)

func TestConcurrentRequests(t *testing.T) {
	start := time.Now()

	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go func(i int) {
			resp, err := http.Get("http://localhost:8000/transactions")
			if err != nil {
				t.Errorf("Request %d error: %v", i, err)
				ch <- "fail"
				return
			}
			resp.Body.Close()
			ch <- "ok"
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-ch
	}

	elapsed := time.Since(start)

	if elapsed.Seconds() > 5 {
		t.Errorf("Expected concurrent handling, but took too long: %.2f seconds", elapsed.Seconds())
	} else {
		t.Logf("Test passed, total time: %.2f seconds", elapsed.Seconds())
	}
}
