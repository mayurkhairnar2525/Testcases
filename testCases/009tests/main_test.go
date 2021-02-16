package main

import "testing"

func TestSum (t *testing.T) {
	var texts = []struct {
		expected int
		numbers  []int
	}{
		{150, []int{10, 20, 30, 40, 50}},
		{230, []int{10, 20, 50, 70, 80}},

	}
		for _, test := range texts{
		if output := Sum(test.numbers); output!=test.expected{
		t.Error("Test Failed: {} inputted, {} expected, received: {}", test.numbers, test.expected, output)}
	}
	}
