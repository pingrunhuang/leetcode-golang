package leetcode

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	a := []int{2, 7, 11, 15}
	b := 9
	result := twoSum(a, b)
	expected := []int{0, 1}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("%v != %v", result[i], expected[i])
		}
	}
}
