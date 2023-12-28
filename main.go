package main

import (
	"fmt"

	"github.com/devzcraft/go-concurrency-patterns/concurrency"
)

func main() {
	ch1 := make(chan any, 10)
	ch2 := make(chan any, 20)
	ch3 := make(chan any, 30)
	channels := []<-chan any{ch1, ch2, ch3}

	sink := concurrency.OrWithChannel(channels)

	for i := 0; i < 35; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		case ch3 <- i:
		}
	}
	close(ch1)
	close(ch2)
	close(ch3)

	for v := range sink {
		fmt.Printf("%+v\n", v)
	}
}
