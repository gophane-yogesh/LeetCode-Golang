package main

import (
	"fmt"
	"sort"
)

func main() {

	input := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}

	res := sortTheStudents(input, 2)
	fmt.Println(res)

}

func sortTheStudents(score [][]int, k int) [][]int {

	n := len(score)
	res := make(map[int]int, n)

	ans := make([][]int, n)
	t := make([]int, 0, len(res))

	for i := 0; i < n; i++ {

		res[score[i][k]] = i
		t = append(t, score[i][k])

	}

	sort.Sort(sort.Reverse(sort.IntSlice(t)))

	for i := 0; i < n; i++ {

		ans[i] = score[res[t[i]]]

	}
	return ans

}
