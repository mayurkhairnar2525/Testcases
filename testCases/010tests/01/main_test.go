package main

import "testing"

func TestAdd (t *testing.T) {
	var texts =[] struct{
		value1 int
		value2 int
		expected int
	}{
		{2,2,4},
		{5,10,15},
	}
	for _,test :=range texts{
		if output := Add(test.value1,test.value2); output != test.expected{
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.value1,test.value2, test.expected, output)		}
	}
}

func TestSub (t *testing.T) {
	var texts =[] struct{
		value3 int
		value4 	int
		expected int
	}{
		{10,2,8},
		{10,5,5},
	}
	for _,test :=range texts{
		if output := Sub(test.value3,test.value4); output != test.expected{
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.value3,test.value4, test.expected, output)		}
	}
}

