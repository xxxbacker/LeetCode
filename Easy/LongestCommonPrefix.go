package main

/*
Input: strs = ["flower","flow","flight"]
Output: "fl"

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.
*/

import "fmt"

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return "err"
	}

	firstString := strs[0]

	lenFirstString := len(firstString)

	commonPrefix := ""
	for i := 0; i < lenFirstString; i++ {
		firstStringChar := string(firstString[i])
		match := true
		for j := 1; j < len(strs); j++ {
			if (len(strs[j]) - 1) < i {
				match = false
				break
			}

			if string(strs[j][i]) != firstStringChar {
				match = false
				break
			}

		}

		if match {
			commonPrefix += firstStringChar
		} else {
			break
		}
	}

	return commonPrefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"fan", "fat", "fame"}))
	fmt.Println(longestCommonPrefix([]string{"bat", "van", "cat"}))
}
