package main

import "fmt"

// make(chan string)
// chan<-
// <-chan

func main() {
	buffer := make(chan string)
	done := make(chan bool)

	go set(buffer, done)

	fmt.Println(<-buffer)
	<-done
}

func set(buf chan<- string, done chan<- bool) {
	buf <- "go"
	done <- true
}
