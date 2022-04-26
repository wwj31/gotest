package main

import (
	"fmt"
	"math/rand"
	"time"
)

//go build -race
func main() {
	ch1 := make(chan int, 10)
	var val = 123
	go func() {
		<-ch1
		for {
			val = rand.Int()
			fmt.Println(len(ch1))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- 0
			time.Sleep(time.Second)
		}
	}()

	_ = val
	select {}
}
