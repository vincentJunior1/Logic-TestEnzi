package main

import (
	"testing"
)

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

func TestAdditionEvent(t *testing.T) {
	test := []struct {
		TestCase []int
		Expected int
	}{
		{[]int{1, 10, 13, 2, 41, 55, 20}, 32},
		{[]int{1, 21, 22, 23, 24, 25, 26}, 72},
		{[]int{1022, 1021, 1023, 2041, 2341}, 1022},
	}

	for i, tt := range test {
		t.Run("", func(t *testing.T) {
			result := AdditionEven(tt.TestCase)
			if result != tt.Expected {
				t.Error(result)
				t.Errorf("Failed running testcase %v", tt.TestCase)
			} else {
				t.Logf("Test case %d Success", i)
			}
		})
	}
}

func TestMaximumAndMinimumNumber(t *testing.T) {
	test := []struct {
		TestCase []int
		Expected []int
	}{
		{[]int{1, 9, 10, 2, 5, 76, 2, 4}, []int{1, 76}},
		{[]int{502, 201, 222, 455, 92, 104}, []int{104, 502}},
		{[]int{1022, 1021, 1023, 2041, 2341}, []int{1021, 2341}},
	}

	for i, tt := range test {
		t.Run("", func(t *testing.T) {
			min, max := MaximumAndMinimumNumber(tt.TestCase)
			if min != tt.Expected[0] && max != tt.Expected[1] {
				t.Errorf("Failed running testcase %v", tt.TestCase)
			} else {
				t.Logf("Test case %d Success", i)
			}
		})
	}
}
