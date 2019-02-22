package leetcode

import (
	"sort"
)

func findPairs(nums []int, k int) int {

	if k < 0 {
		return 0
	}
	memo := make(map[int]int)
	for _, v := range nums {
		memo[v]++
	}

	ans := 0
	if k == 0 {
		for _, tmepV := range memo {
			if tmepV > 1 {
				ans++
			}
		}
		return ans
	}
	for tempK := range memo {
		if memo[tempK-k] > 0 {
			ans++
		}
	}
	return ans
}

func findPairs2(nums []int, k int) int {
	// increasing order
	sort.Ints(nums)
	n := len(nums)
	ans := 0
	left := 0
	right := 1
	for left < n && right < n {
		leftNum := nums[left]
		rightNum := nums[right]

		if right-left < k {
			right++
		}
		if right-left > k {
			left++
		}

		if rightNum-leftNum == k {
			ans++
			// remove the duplicate
			for left < n && nums[left] == leftNum {
				left++
			}
			for right < n && nums[right] == rightNum {
				right++
			}
		}
		if right == left {
			right++
		}
	}
	return ans
}
