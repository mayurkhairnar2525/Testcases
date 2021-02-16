package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(1, 2)
	assert.NotNil(t, total, "The total should be not nil")
	assert.Equal(t, 3, total, "Expecting 3")

}

func TestSub(t *testing.T) {
	total := Sub(10,5)
	assert.NotNil(t,total,"The total should be not nil")
	assert.Equal(t,5,total,"Expecting 5")
}