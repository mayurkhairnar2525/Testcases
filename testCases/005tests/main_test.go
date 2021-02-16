package main

import "testing"

func TestConverter (t *testing.T) {
	input :="Thisisasamplestring "
	Expected :="Thisisasamplestring"
	//input :="This_is_a_sample_string "
	//Expected := "This is a sample string"
	actual ,err :=(Converter(input))
	if err != nil {
		t.Errorf("expected no error,but got %v",err)
	}
	if actual != Expected{
		t.Errorf("expected output to be %s ,but got %s ",Expected,actual)
	}
}
