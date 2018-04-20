package main

import "fmt"
import "strings"

// Given a string, find the length of the longest substring without repeating characters.
// Examples:
// Given "abcabcbb", the answer is "abc", which the length is 3.
// Given "bbbbb", the answer is "b", with the length of 1.
// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	cur := 0
	startCur := 0
	maxLength := 0
	subStr := ""
	for cur < len(s) && startCur < len(s) {
		if strings.Contains(subStr, string(s[cur])) {
			// 坑：当遇到一个字符第一次重复出现在最近刚生成的子串的时候，要注意这时候startCur的值，应该是原先
			startCur = startCur + strings.Index(subStr, string(s[cur])) + 1
			cur = startCur + 1
			subStr = string(s[startCur])
		} else {
			subStr = subStr + string(s[cur])
			cur++
		}
		// 坑：应该每次都计算这个最大子串是否是最大, 这样避免像"dvdf"忽略了"vdf"的可能性
		if maxLength < len(subStr) {
			maxLength = len(subStr)
		}
	}
	return maxLength
}

// 如果采用hashmap，应该可以更加优化
func main() {
	str1 := "abcabcbb"
	result1 := lengthOfLongestSubstring(str1)
	fmt.Println(result1)
	if result1 == 3 {
		fmt.Printf("Test case #3 %s passed\n", str1)
	} else {
		fmt.Printf("Test case #3 %s is not passed\n", str1)
	}

	str2 := "bbbbbb"
	result2 := lengthOfLongestSubstring(str2)
	fmt.Println(result2)
	if result1 == 1 {
		fmt.Printf("Test case #3 %s passed\n", str2)
	} else {
		fmt.Printf("Test case #3 %s is not passed\n", str2)
	}

	str3 := "dvdf"
	result3 := lengthOfLongestSubstring(str3)
	fmt.Println(result3)
	if result3 == 1 {
		fmt.Printf("Test case #3 %s passed\n", str3)
	} else {
		fmt.Printf("Test case #3 %s is not passed\n", str3)
	}
}
