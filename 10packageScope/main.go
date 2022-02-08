package main

import "fmt"

var score = 99.5

func main() {
	//num can't be accessed by other file as it is not in the global scope of package main
	num := 40
	sayHello("tina")
	for _, v := range points {
		fmt.Println(v)
	}
	showScore()
}
