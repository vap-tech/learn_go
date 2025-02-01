// echo v2.0
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for i, arg := range os.Args[0:] {
		fmt.Println(i, arg)
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
}
