package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	time.Sleep(500 * time.Millisecond)
	fmt.Println(time.Now().Sub(a).Seconds())

	// sec with decimal to duration
	n := 12.34
	fmt.Println(time.Duration(n*1000000000) * time.Nanosecond)
	fmt.Println(time.Duration(n*1e9) * time.Nanosecond)
}
