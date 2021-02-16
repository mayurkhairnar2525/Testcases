package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//Test Table
func TestSum (t *testing.T) {
	if sum(5,5)!=10 {
		t.Errorf("Expcetd 10")
	}
}
func TestTableSum (t *testing.T) {
	var texts = []struct {
		input    int
		inputt   int
		expected int
	}{
		{10, 10, 20},
		{50, 10, 60},
	}
	for _, test := range texts {
		if output := sum(test.input, test.inputt); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}

//With assert package
func TestSumAssert (t *testing.T) {
	total := sum (10,15)
	assert.NotNil(t,total,"The total should be not nil")
	assert.Equal(t,25,total,"Expecting 5")
}
