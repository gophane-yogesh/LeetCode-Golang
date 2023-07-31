package main

import "fmt"

func main() {

	input := []string{"floower", "flow", "flight"}

	res := longestCommonPrefix(input)

	fmt.Print(res)

}

func longestCommonPrefix(strs []string) string {
	res := ""

	slen := 0

	for _, val := range strs{
		slen = min(slen,len(val))
	}

	for i:=0; i<slen; i++{
		for j:=0; j<len(strs);j++{
			if strs[j][i] == strs[j+1][i]{
				continue
			}
 

		}
		res:= strs[]
	}


	return res

}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}