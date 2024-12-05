package main

import "testing"

func TestMaximumProfitStock(t *testing.T) {
	test := []struct {
		TestCase []int
		Expected int
	}{
		{[]int{2, 5, 6, 1, 3}, 4},
		{[]int{10, 2, 1, 4, 6}, 5},
		{[]int{8, 2, 5, 1, 5, 7}, 6},
	}

	for i, tt := range test {
		t.Run("", func(t *testing.T) {
			result := MaximalProfitStock(tt.TestCase)
			if result != tt.Expected {
				t.Errorf("Failed running testcase %v", tt.TestCase)
			} else {
				t.Logf("Test case %d Success", i)
			}
		})
	}
}
