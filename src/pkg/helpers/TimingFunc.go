package helpers

import (
	"fmt"
	"net/http"
	"time"
)

// Elapsed ..
func Elapsed(r *http.Request) func() {
	start := time.Now()
	return func() {
		elapsed := float64(time.Now().Sub(start)) / float64(time.Millisecond)
		msg := fmt.Sprintf("ROUTE: %s - %fms", r.URL.Path, elapsed)
		fmt.Println(msg)
		// if elapsed > float64(time.Millisecond/1000) {
		// 	go ErrorStore(r, "TIMING", "SERVER", msg, "", "", 0)
		// }
	}
}
