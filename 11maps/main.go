package main

import "fmt"

func main() {
	//key:strings,values:float64 in the below case
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["salad"])

	//looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		111111: "tina",
		222222: "john",
		333333: "mark",
		444444: "nia",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[222222])

	//update
	phonebook[111111] = "blah"
	fmt.Println(phonebook)
}
