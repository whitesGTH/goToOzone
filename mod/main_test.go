package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOnesAfterRemoveItem(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(maxOnesAfterRemoveItem([]byte{0, 0}), uint(0))
	assert.Equal(maxOnesAfterRemoveItem([]byte{0, 1}), uint(1))
	assert.Equal(maxOnesAfterRemoveItem([]byte{1, 0}), uint(1))
	assert.Equal(maxOnesAfterRemoveItem([]byte{1, 1}), uint(1))
	assert.Equal(maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1}), uint(4))
	assert.Equal(maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1, 0, 1, 1, 1}), uint(5))
	assert.Equal(maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1, 0, 1, 1, 1, 0}), uint(5))
}
