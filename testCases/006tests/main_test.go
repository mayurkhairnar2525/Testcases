package main

import (
"testing"
)

func TestReverseString(t *testing.T) {
	if ReverseString("") != "" {
		t.Fail()
	}



    if ReverseString("hello") != "olleh" {
   t.Fail()
    }
}
