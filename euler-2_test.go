package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstIndex(t *testing.T) {
	res := Fib(0)
	assert.Equal(t, 1, res)
}

func TestSecondIndex(t *testing.T) {
	res := Fib(1)
	assert.Equal(t, 2, res)
}

func TestLastIndex(t *testing.T) {
	res := Fib(9)
	assert.Equal(t, 89, res)
}
