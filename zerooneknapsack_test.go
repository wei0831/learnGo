package learngo

import "testing"

type knapsackInput struct {
	w  int64
	wt []int64
	vt []int64
}

type knapsackTestCase struct {
	knapsackInput
	expected int64
}

func TestKnapSack(t *testing.T) {
	cases := []knapsackTestCase{
		{
			knapsackInput{50,
				[]int64{10, 20, 30},
				[]int64{60, 100, 120},
			},
			220,
		},
		{
			knapsackInput{7,
				[]int64{1, 3, 4, 5},
				[]int64{1, 4, 5, 7},
			},
			9,
		},
		{
			knapsackInput{2,
				[]int64{},
				[]int64{},
			},
			0,
		},
		{
			knapsackInput{0,
				[]int64{},
				[]int64{},
			},
			0,
		},
		{
			knapsackInput{0,
				[]int64{1, 2, 3},
				[]int64{1, 2, 3},
			},
			0,
		},
	}

	for _, c := range cases {
		r := KnapSack(c.knapsackInput.w, c.knapsackInput.wt, c.knapsackInput.vt)
		if r != c.expected {
			t.Errorf("knapSack(%+v) == %v, expected %v", c.knapsackInput, r, c.expected)
		}
	}
}
