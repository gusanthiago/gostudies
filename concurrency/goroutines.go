// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func log(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	log("direct")
	go log("go")
	go log("foo")

	time.Sleep(time.Second + 1)
	fmt.Println("done")
}
