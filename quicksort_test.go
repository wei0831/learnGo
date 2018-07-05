package learngo

import "testing"
import "sort"

type quickSortTestCase struct {
	input    []int
	expected []int
}

func TestQuickSort(t *testing.T) {
	cases := []quickSortTestCase{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{9}, []int{9}},
		{[]int{9, 8}, []int{8, 9}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 3, 1, 2, 4}, []int{1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		r := make([]int, len(c.input))
		copy(r, c.input)
		QuickSort(r)
		if len(c.expected) != len(r) {
			t.Errorf("QuickSort(%+v) == %+v, expected %+v", c.input, r, c.expected)
			continue
		}
		for i := range c.expected {
			if c.expected[i] != r[i] {
				t.Errorf("QuickSort(%+v) == %+v, expected %+v", c.input, r, c.expected)
				break
			}
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for n := 0; n < b.N; n = n + 1 {
		input := []int{5, 12, 57, 35, 68, 14, 78, 34, 93, 64, 2, 99}
		QuickSort(input)
	}
}

func BenchmarkGoSort(b *testing.B) {
	for n := 0; n < b.N; n = n + 1 {
		input := []int{5, 12, 57, 35, 68, 14, 78, 34, 93, 64, 2, 99}
		sort.Ints(input)
	}
}
