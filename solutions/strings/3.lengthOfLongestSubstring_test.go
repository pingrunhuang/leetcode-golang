package strings

import (
	"fmt"
	"testing"
)

// 如果采用hashmap，应该可以更加优化
func TestLengthOfLongestSubString(t *testing.T) {
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
