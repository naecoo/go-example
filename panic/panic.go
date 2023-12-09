package main

import "os"

func main() {

	panic("a problem")

	_, err := os.ReadFile("./tmp/file")
	if err != nil {
		panic(err)
	}
}
