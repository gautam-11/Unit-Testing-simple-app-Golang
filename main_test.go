package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	total := Add(1, 3)
	assert.NotNil(t, total, "The `total` should not be `nil`")
	assert.Equal(t, 4, total, "Expecting `4`")
}

func TestSubtract(t *testing.T) {
	difference := Subtract(1, 3)
	assert.NotNil(t, difference, "The `difference` should not be `nil`")
	assert.Equal(t, -2, difference, "Expecting `-2`")
}

func TestMultiply(t *testing.T) {

	product := Multiply(4, -2)
	assert.NotNil(t, product, "The `product` should not be nil")
	assert.Equal(t, -8, product, "Expecting `-8` ")

}

func TestDivide(t *testing.T) {

	division, err := Divide(4, 2)
	assert.NotNil(t, division, "The `division` should not be nil")
	assert.Equal(t, 2, division, "Expecting `2`")

	division, err = Divide(4, 0)
	assert.NotNil(t, err, "The `err` should not be nil")
	assert.Equal(t, "Division by zero exception", err.Error(), "Expecting Division by zero exception")
	assert.Equal(t, 1<<16-1, division, "Expecting `1<<16-1`")

}
