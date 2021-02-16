package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//Simple test
func TestCalculate(t *testing.T)  {
	if Calculate(2)!=4 {
		t.Errorf("Expcetd 4")
	}
}
//Test Table
func TestTableCalculate(t *testing.T)  {
	var texts = [] struct {
		input int
		expected int
		}{

		{2,4},
		{5,7},
		{9,11},
		{500,502},
		{1000,1002},
	}
	for _,test := range texts{
		if output :=Calculate(test.input);output!=test.expected{
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)		}
	}

}

//using t.Run

func TestCalculateInt (t *testing.T) {
	var texts = []struct{
		name string
		numbers int
		expected int
	}{
		{"First one",20,22},
		{"Second one",10,12},
		{"Third one",25,27},
		{"Fourth one",8,10},
		{"Five one",13,15},

	}
	for _, tc := range texts {
		t.Run(tc.name, func(t *testing.T) {
			s := Calculate(tc.numbers)
			if s != tc.expected {
				t.Fatalf("sum of %v should be %v; got %v", tc.name, tc.expected, s)
			}
		})
	}
}

//with assert
func TestAdd (t *testing.T){
	total :=Calculate(5)
	assert.NotNil(t, total, "The total should be not nil")
	assert.Equal(t, 7, total, "Expecting 3")
    assert.NotEqual(t,12,13,"They should be not be Equal")
}