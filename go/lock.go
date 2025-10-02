package main
import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool)
	for i:=0; i<7; i++ {
		go worker(i, lock)
	}
	time.Sleep(10*time.Second)
}

func worker(id int, lock chan<- bool) {
	fmt.Println(”%d wants the lock \n”, id)
	lock <− true
	fmt.Println(”%d has the lock \n”, id)
	time.Sleep(500 ∗ time.Millisecond)
	fmt.Println(”%d is releasing the lock \n”, id)
	<−lock
}