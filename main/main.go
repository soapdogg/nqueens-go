package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano() / int64(time.Millisecond)
	orchestrate(9)
	end := time.Now().UnixNano() / int64(time.Millisecond)

	elapsedMillis := end - start
	fmt.Println("Time elapsed (millis): ", elapsedMillis)
}
