package main
import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1) 
	go send(ch2)

	select {
		case i1 := <-ch1: fmt.Printf("first call %d \n",i1)
		case i2 := <-ch2: fmt.Printf("second call %d \n",i2) 
	//default: fmt.Println("I dont block")
	}
}

func send(ch chan<- int) {
	ch <- 1
}