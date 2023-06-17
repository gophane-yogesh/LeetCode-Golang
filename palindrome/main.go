package main

import "fmt"

func main() {

	// call function with static parameters

	res := isPalindrome(121)

	fmt.Print(res)

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	num := x // keep copy of x
	rev := 0

	for x != 0 {

		rem := x % 10
		x = x / 10

		rev = rev*10 + rem

	}

	return rev == num

}
