package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	result := Add(1,2)
	assert.Equal(t, result, 3)
}