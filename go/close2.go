package main
import "fmt"

func main() {
	ch := make(chan int)
	go send(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) // prints 0 because ch is closed
}

func send(ch chan<- int) {
	ch <- 1; close(ch) 
}