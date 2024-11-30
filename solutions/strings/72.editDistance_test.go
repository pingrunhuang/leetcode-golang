package strings

import "testing"

type testCase struct {
	test []string
	ans  int
}

func TestMinDis(t *testing.T) {
	testCases := []testCase{
		{[]string{"horse", "ros"}, 3},
		{[]string{"intention", "execution"}, 5},
	}
	for _, v := range testCases {
		result := minDistance(v.test[0], v.test[1])
		if v.ans != result {
			t.Error("Expected: ", v.ans, "Got:", result)
		}
	}

}
