package main

import (
	"fmt"
	"runtime"
)

func main() {
	var v uint
	var arr []int8

	for {
		fmt.Print("How many MB to allocate? ")
		_, err := fmt.Scanln(&v)
		if err != nil {
			panic(err)
		}

		// So gc doesn't get rid of arr
		for i := range arr {
			arr[i] = 2
		}
		fmt.Printf("Previously allocated: %v bytes.\n", len(arr))
		arr = make([]int8, 1024*1024*v)
		for i := range arr {
			arr[i] = 1
		}
		fmt.Printf("New allocated %v bytes.\n", len(arr))
		runtime.GC()
	}

}
