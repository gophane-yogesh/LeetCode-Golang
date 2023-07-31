package main

import (
	"fmt"
)

func main() {

	input := []int{0, 2, 1, 5, 3, 4}

	res := getConcatenation(input)

	fmt.Print(res)

}

// TC : O(N)  SC : O(N)

func buildArray(nums []int) []int {

	var ans = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}

	return ans

}

func getConcatenation(nums []int) []int {
	n := len(nums)

	var ans = make([]int, 2*n)

	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans

}
