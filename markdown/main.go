package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Not enough arguments")
	}

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	data = blackfriday.MarkdownCommon(data)

	fmt.Println(string(data))
}
