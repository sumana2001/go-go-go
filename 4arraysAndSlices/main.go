package main

import "fmt"

func main() {
	//arrays
	var ages1 [3]int = [3]int{20, 25, 30}

	//shorthand
	var ages2 = [3]int{20, 25, 30}

	//even more shorthand
	names := [4]string{"john", "tina", "mark", "nia"}
	fmt.Println(ages1, ages2, "Length of ages1:", len(ages1))
	fmt.Println(names, len(names))

	//update array
	names[1] = "hehe"
	fmt.Println(names, len(names))

	//slices (variable length) (use arrays under the hood)
	var scores = []int{100, 50, 60}
	//update value
	scores[2] = 25
	//add new value
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	//slice ranges
	//from 1 upto  3 (include 1 exclude 3)
	rangeOne := names[1:3]
	//from 2 till end (include 2)
	rangeTwo := names[2:]
	//from start upto 3 (exclude 3)
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)
	//these slice ranges also become individual slices
	rangeOne = append(rangeOne, "blah")
	fmt.Println(rangeOne)
}
