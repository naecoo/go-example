package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// id
		wg.Add(1)

		i := i

		go func() {
			// worker完毕后，发送信号给wg
			defer wg.Done()
			worker(i)
		}()
	}

	// 等待所有worker完成
	wg.Wait()

}
