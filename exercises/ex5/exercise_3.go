package main

import (
	"fmt"
	"sync"
)

func f(fut chan<- int, p1 int, p2 int) {
	var once sync.Once
	once.Do(func() {
		fut <- p1 + p2
	})
	close(fut)
}

func main() {
	ch := make(chan int)
	go f(ch, 1, 2)

	for i := 0; i < 5; i++ {
		val, ok := <-ch
		if ok {
			fmt.Println(val)
		}
	}
}
