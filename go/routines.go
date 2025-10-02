package main
import (
	"fmt"
	"net/http"
)

// go checklink(link, c)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links { 
		go checklink(link, c) 
	}
	for i := 0; i < len(links); i++ { 
		fmt.Println(<-c) 
	}
}

func checklink(link string, c chan<- string) {
	_, err := http.Get(link) //pairs as language builtins

	if err != nil {
		c <- link + " is down!"
		return
	}
	c <- link + " is up!"
}