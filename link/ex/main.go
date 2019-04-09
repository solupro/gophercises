package main

import (
	"fmt"
	"github.com/solupro/gophercises/link"
	"os"
)

func main() {
	f, err := os.Open("./ex2.html")
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
