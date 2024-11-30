package strings

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (39.44%)
 * Total Accepted:    150.9K
 * Total Submissions: 382.7K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome.
 *
 * Return all possible palindrome partitioning of s.
 *
 * Example:
 *
 *
 * Input: "aab"
 * Output:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 * Back-tracking
 */

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	return s == reverseString(s)
}

func _partition(str string, start int, tempRes []string, res *[][]string) {
	if start >= len(str) {
		newElement := make([]string, len(tempRes))
		copy(newElement, tempRes)
		*res = append(*res, newElement)
		return
	}

	for i := start; i < len(str); i++ {
		// notice substring is not including the right
		if isPalindrome(str[start : i+1]) {
			if start == i {
				tempRes = append(tempRes, string(str[i]))
			} else {
				tempRes = append(tempRes, str[start:i+1])
			}
			_partition(str, i+1, tempRes, res)
			tempRes = tempRes[:len(tempRes)-1] // remove the last added element
		}
	}
}

func partition(s string) [][]string {
	result := [][]string{}
	tempRes := []string{}
	_partition(s, 0, tempRes, &result)
	return result
}
