package main

import (
	"fmt"
	"sort"
)

func main() {

	nums1 := []int{0,1,2,2,3,0,4,2} // 1, 2, 2,3, 0, 0 ,0
	// nums2 := []int{2, 5, 6}          // 3, 5, 6
	// x := 3
	res := removeElement(nums1, 2)
	fmt.Print(res)

}

func merge(nums1 []int, m int, nums2 []int, n int) {

	j := 0
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[j]
		j++
	}

	sort.Ints(nums1)

}

func removeElement(nums []int, val int) int {
	k := 0

	for i := 0; i < len(nums); i++ {

		if val == nums[i] {
			nums[i] = 51
			k++
		}

	}
	sort.Ints(nums)
	fmt.Print(nums)
	return len(nums)-k

}

// check both index of arrays
// if num1 is less then dont swap
// else swap the numbers and increase to index  + 1
