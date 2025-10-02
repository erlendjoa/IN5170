package main
import (
	"os"
	"fmt"
	"sync"
	"compress/gzip"
)

// sync.WaitGroup
// wg.Add(1)
// wg.Done()
// wg.Wait()

func main() {
	var wg sync.WaitGroup
	var i int
	var file string
	
	for i, file = range os.Args[1:] {
		wg.Add(1) // add before asynchronous call!

		go func(file string) { // anonymous function
			compress(file)
			wg.Done()
		} (file)
	}
	wg.Wait()
	fmt.Printf("compressed %d files \n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	defer gzout.Close()
	
	return nil
}