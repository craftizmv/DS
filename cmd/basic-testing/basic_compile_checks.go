package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	mapWriteCheck()
	context.Background()
}

func mapWriteCheck() {
	var m map[string]string
	m = make(map[string]string)
	m["a"] = "b"
	//value, ok := m["a"]
	fmt.Println(m["a"])
}

func chanCheck() {
	c := make(chan int)

	go func() {
		c <- 1
	}()

	time.Sleep(time.Microsecond * 1000)

	select {
	case <-c:
		fmt.Println("1")
	default:
		fmt.Println("2")
	}
}
