package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxOnesAfterRemoveItem(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(uint(0), maxOnesAfterRemoveItem([]byte{0, 0}))
	assert.Equal(uint(1), maxOnesAfterRemoveItem([]byte{0, 1}))
	assert.Equal(uint(1), maxOnesAfterRemoveItem([]byte{1, 0}))
	assert.Equal(uint(1), maxOnesAfterRemoveItem([]byte{1, 1}))
	assert.Equal(uint(4), maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1}))
	assert.Equal(uint(5), maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1, 0, 1, 1, 1}))
	assert.Equal(uint(5), maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1, 0, 1, 1, 1, 0}))
}
