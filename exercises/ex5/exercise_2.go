package main

import "fmt"

type OP int

const (
	ADD   OP = 0
	INC   OP = 1
	STORE OP = 2
	DUAL  OP = 3
	SNGL  OP = 4
)

type Msg struct {
	op  OP
	p1  int
	p2  int
	ret chan int
}

type State struct {
	mode OP
}

func loop1(ch chan Msg, state State) {
	var memory int
	for {
		msg := <-ch

		if state.mode == DUAL {
			switch msg.op {
			case ADD:
				msg.ret <- msg.p1 + msg.p2
			case SNGL:
				state.mode = SNGL
			}
		} else if state.mode == SNGL {
			switch msg.op {
			case STORE:
				memory = msg.p1
			case INC:
				msg.ret <- memory + msg.p1
			case DUAL:
				state.mode = DUAL
			}
		}
	}

}

func main() {
	input := make(chan Msg)
	go loop1(input, State{mode: SNGL})
	res := make(chan int)
	input <- Msg{STORE, 2, 0, res}
	input <- Msg{INC, 5, 0, res}
	fmt.Println(<-res) // should print 7
}

/* LEARNING OUTCOMES */

// go loop1(input, State{mode: SNGL})

// for {
//		msg := <-ch
//		... }

// type state Struct { mode OP }
// st := Struct{mode: SNGL}
