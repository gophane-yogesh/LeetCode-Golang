
package main

import "fmt"

func main() {

	fmt.Print(twoSum([]int{2, 7, 11, 15}, 9))

}

func twoSum(nums []int, target int) []int {
  
  	// Optimal Solution  TC : O(N) SC : O(N)

	var mp = make(map[int]int)

	var res []int

	for index, currNum := range nums {

		revNum := target - currNum

		if val, ok := mp[revNum]; ok {
			res = append(res, index, val)

		}
		mp[currNum] = index

	}

	return res

	// Brute Force Solution TC : O(N^2) SC : O(1)
	// 	n := len(nums) // length of nums slice
	// 	var res []int
	// 	for i := 0; i < n; i++ {
	// 		for j := i + 1; j < n; j++ {
	// 			if nums[i] + nums[j] == target {
	// 				res = append(res, i, j) // stores index
	// 				break
	// 			}
	// 		}
	// 	}
	// 	return res
	// }



}
