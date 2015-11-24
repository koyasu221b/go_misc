package main

import "fmt"

var author = "hi"

var Version = "0.0.1"

func main() {
	var gretting string = "Hello world!"
	var a, b, c int = 1, 2, 3
	fmt.Println(gretting, a, b, c)

	var name = "Jeremy Saenz"
	var d, e, f = 1, 2.0, "Three"
	fmt.Println(name, d, e, f)

	course := "Essential Go"
	x, y, z := 1, 2, 3
	fmt.Println(course, x, y, z)
}
