package main

import "fmt"

func main() {
	mybill := newBill("Tina's bill")
	fmt.Println(mybill.format())
}
