package main

import "testing"

func TestHello(t *testing.T) {
	Got := Hello()
	Expected := "Hello World"
	if Got!=Expected {
		t.Errorf("get %q want %q",Got,Expected)
	}
}





/*
If there is string in the function then 1st declared variable like we did "got" and "Expected".
if you directly start with if statement like we previously did for int (in 0001test and 0002test)
so for string don't do that because it not works.
 */