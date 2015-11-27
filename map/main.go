package main

import "fmt"

func main() {
	age := make(map[string]int)

	age["hello"] = 20
	age["world"] = 30
	fmt.Println(age)

	fmt.Println("hello is :", age["hello"])

	delete(age, "world")
	fmt.Println(age)

	m := map[string]int{
		"hello": 21,
		"world": 22,
	}
	fmt.Println(m)

	for n, a := range m {
		fmt.Printf("%v is %v\n", n, a)
	}
}
