package main

import "fmt"

func main() {
	mybill := newBill("Tina's bill")
	mybill.updateTip(2)
	mybill.addItem("donut", 1.99)
	fmt.Println(mybill.format())
}
