package substring

import (
	"fmt"
	"strings"
)

// too slow
func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	r := []rune(s)
    for range r {
		tempLength := 0
		substring := ""
		for _, char := range r {
			if strings.Contains(substring, string(char)) {
				if tempLength > maxLength {
					maxLength = tempLength
				}
				tempLength = 0
				substring = ""
			} else {
				substring += string(char)
				tempLength++
			}
		}
		r = r[1:]
		if tempLength > maxLength {
			maxLength = tempLength
		}
		tempLength = 0
	}
	return maxLength
}

func lengthOfLongestSubstringRedux(s string) int {
	// create a map of the character to its position (length)
	m := make(map[byte]int)

	head, tail, res := 0, -1, 0

	// maximum length has to decrease as we move through the string
	for head < len(s) {
		// j 
		if j, ok := m[s[head]]; ok {
			fmt.Println(m)
			fmt.Println(j)
			fmt.Println(string(s[head]))
			tail = maxInt(tail, j)
		}
		m[s[head]] = head
		res = maxInt(res, head-tail)
		head++
	}
	return res
}

func maxInt(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}