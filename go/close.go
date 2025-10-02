package main
import (
	"fmt"
	"time"
)

// close(ch)

func main() {
	ch := make(chan int)
	go send(ch)
	for {
		i, ok := <-ch
		if !ok {break}
		fmt.Println(i) 
	} 
	time.Sleep(1 * time.Second)
}

func send(ch chan<- int) {
	ch <- 1; ch <- 2;
	close(ch) 
}
