package main

import (
	"io/ioutil"
	"testing"
)

//We can do it using both methods

/*
type AddResult struct{
	x int
	y int
	expected int
}
var addResult= [] AddResult{
	{1,1,2},
	{100,105,205},
	{500,107,607},
	{405,5,410},
	{650,50,700},
}
func TestAdd (t *testing.T) {
	for _,test := range addResult{
		result := Add(test.x,test.y)
		if result != test.expected{
			t.Errorf("Expected resuly not given")

//Instead of t.Errorf also we can use t.Fatal. Like,
//t.Fatal("expected result not given")
		}
	}
}
 */
func TestAdd (t *testing.T) {
	var texts =[] struct{
		x int
		y int
		expected int
	}{
		{2,2,4},
		{5,10,15},
	}
	for _,test :=range texts{
		if output := Add(test.x,test.y); output != test.expected{
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.x,test.y, test.expected, output)		}
		}
	}



//TeestData/Read Test data

func TestReadFile(t *testing.T) {
	data,err :=ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not Open")
	}
	if string (data)!="hello world "{
		t.Fatal("string contain does not match")
	}
}


