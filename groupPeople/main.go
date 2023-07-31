package main

import "fmt"

func main() {

	input := []int{4, 5, 6, 7, 0, 2, 1, 3}

	np := input[2:4]
	go restoreString()

	// select {}

	fmt.Print(cap(np))

}

// func groupThePeople(groupSizes []int) [][]int {
//     var res  = make(string[int],10)

// }

func restoreString() {
	fmt.Print("hello Anjali")

}
