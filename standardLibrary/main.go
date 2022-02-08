package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//strings package
	greeting := "Hey there!"
	fmt.Println(strings.Contains(greeting, "Hey"))
	fmt.Println(strings.ReplaceAll(greeting, "e", "E"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ey"))
	fmt.Println(strings.Split(greeting, " "))

	//original value is still unchanges
	fmt.Println(greeting)

	//sort package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)

	index1 := sort.SearchInts(ages, 30)
	//if not found returns len
	index2 := sort.SearchInts(ages, 90)
	fmt.Println(index1, index2)

	names := []string{"john", "tina", "mark", "nia"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "nia"))
}
