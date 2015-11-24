package main

import "fmt"
import "strings"

func main() {
	fmt.Println(double(5))

	first, _ := parseName("jermey Saenz")
	fmt.Println(first)

	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	greet(first)
}

func double(n int) int {
	return n + n
}

func parseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}
