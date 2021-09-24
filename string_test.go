package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfString(t *testing.T) {
	assert.True(t, IfString(true, "1", "0") == "1")
	assert.True(t, IfString(false, "1", "0") == "0")
}

func TestStringOr(t *testing.T) {
	assert.True(t, StringOr( "", "0", "") == "0")
}
