package main

import "testing"

func TestInts (t *testing.T) {
	var texts = []struct{
		name string
		numbers [] int
		expected int
	}{
		{"First one",[] int {10,20},30},
		{"Second one",[] int {10,20,30,40,50,},150},
		{"Third one",[] int {1,20,25},46},
		{"Fourth one",[] int {10,0},10},
		{"Five one",[] int {10,20},30},

	}
	for _, tc := range texts {
		t.Run(tc.name, func(t *testing.T) {
			s := Ints(tc.numbers...)
			if s != tc.expected {
				t.Fatalf("sum of %v should be %v; got %v", tc.name, tc.expected, s)
			}
		})
	}
}

