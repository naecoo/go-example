package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// wait ticker
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	// 可以立即消费3个
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 剩余的间隔200ms
	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 塞入5个请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	go func() {
		i := 1
		for range time.Tick(100 * time.Millisecond) {
			burstyRequests <- i
			i++
		}
	}()

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
