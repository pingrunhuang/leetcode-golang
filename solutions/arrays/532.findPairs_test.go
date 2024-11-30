package arrays

import (
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func TestFindPairs(t *testing.T) {
	qs := []question{

		{
			para{[]int{1, 2, 3, 4, 5}, -1},
			ans{0},
		},

		{
			para{[]int{3, 1, 4, 1, 5}, 2},
			ans{2},
		},

		{
			para{[]int{1, 2, 3, 4, 5}, 1},
			ans{4},
		},

		{
			para{[]int{3, 1, 4, 1, 5}, 0},
			ans{1},
		},

		{
			para{[]int{3, 3, 5, 5, 5}, 2},
			ans{1},
		},

		// 如需多个测试，可以复制上方元素。
	}
	for _, q := range qs {
		result := findPairs(q.para.one, q.para.k)
		if q.ans.one != result {
			t.Error("Expected ", q.ans.one, "Got ", result)
		}
	}
}
