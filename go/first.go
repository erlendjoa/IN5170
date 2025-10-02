package main
import "fmt"


type StringStruct struct {X,Y string}

func (s StringStruct) concatenate() string {
	return s.X + " " + s.Y
}

func (s *StringStruct) consolidate() {
	s.X = s.Y
}

func main() {
	var s1 StringStruct = StringStruct{"my", "first"}
	s2 := StringStruct{"go", "program"}

	fmt.Println(s1.concatenate(), s2.concatenate())

	s1.consolidate()
	s2.consolidate()
	fmt.Println(s1.X, s2.X)
}