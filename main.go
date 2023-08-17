package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker: %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	chanel := make(chan int)
	for i := 0; i < 10000; i++ {
		go worker(i, chanel)
	}

	for i := 0; i < 100000; i++ {
		chanel <- i
	}

}
