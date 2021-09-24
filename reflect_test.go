package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameStruct(t *testing.T) {
	type DemoStruct1 struct {

	}
	type DemoStruct2 struct {

	}
	assert.True(t, !IsSameStruct(DemoStruct2{}, DemoStruct1{}))
	assert.True(t, IsSameStruct(DemoStruct1{}, DemoStruct1{}))
}
