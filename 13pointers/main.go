package main

import "fmt"

func updateName(n *string) {
	*n = "tina"
}

func main() {
	name := "john"
	// updateName(name)
	fmt.Println(name)
	fmt.Println("Memory address of name is:", &name)
	m := &name
	fmt.Println(m)
	fmt.Println("Value at memory address:", *m)
	updateName(m)
	fmt.Println(name)
}
