package main

import "fmt"

func main() {

	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
