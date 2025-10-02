package main
import "fmt"

type AbsI interface { 
	Abs() int 
}

type Pair struct { 
	X, Y int 
}
func (p Pair) Abs() int { 
	return p.X + p.Y 
}

type Triple struct { 
	X, Y, Z int 
}
func (t Triple) Abs() int { 
	return t.X + t.Y + t.Z 
}

func main() {
	var a AbsI //must contain something that implements AbsI

	pair := Pair{1,2}
	triple := Triple{3,4,5}
	
	a = &pair // a *Pair is ok
	a = &triple // a *Triple is ok

	fmt.Println(a)
}

