package utils

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type DemoStruct1 struct {
}
type DemoStruct2 struct {
}
type DemoStruct3 struct {
}

func TestIsSameStruct(t *testing.T) {
	assert.True(t, IsSameStruct(DemoStruct1{}, DemoStruct1{}))
	assert.False(t, IsSameStruct(DemoStruct2{}, DemoStruct1{}))
}

func TestIsInstanceIn(t *testing.T) {
	assert.True(t, IsInstanceIn(DemoStruct1{}, ConvertToTypes(DemoStruct1{}, DemoStruct2{})...))
	assert.False(t, IsInstanceIn(DemoStruct3{}, ConvertToTypes(DemoStruct1{}, DemoStruct2{})...))
}

type DemoStructFields struct {
	Id   string
	Name string
}

func TestEachStructField(t *testing.T) {
	EachStructField(DemoStructFields{
		Id:   "1",
		Name: "goal",
	}, func(field reflect.StructField, value reflect.Value) {
		switch field.Name {
		case "Id":
			assert.True(t, value.String() == "1")
		case "Name":
			assert.True(t, value.String() == "goal")
		default:
			assert.Error(t, errors.New("error"))
		}
	})
}

func TestGetTypeKey(t *testing.T) {
	key := GetTypeKey(reflect.TypeOf(DemoStructFields{}))
	assert.True(t, key == "github.com/goal-framework/utils.DemoStructFields")
	assert.True(t, GetTypeKey(reflect.TypeOf(&DemoStructFields{})) == "*github.com/goal-framework/utils.DemoStructFields")
	str := "goal"
	fmt.Println(GetTypeKey(reflect.TypeOf(str)))
	fmt.Println(GetTypeKey(reflect.TypeOf(&str)))
}


func TestWithoutNil(t *testing.T) {
	var (
		demoParam1    = DemoStructFields{Id: "1"}
		demoParam2nil *DemoStruct1
	)
	fmt.Println(WithoutNil(demoParam2nil, demoParam1))
	assert.NotNil(t, WithoutNil(demoParam2nil, demoParam1))
	assert.Nil(t, WithoutNil(demoParam2nil))

	assert.NotNil(t, WithoutNil(demoParam2nil, func() interface{} {
		return demoParam1
	}) != nil)
}

func TestParseStructTag(t *testing.T) {
	rawTag := reflect.StructTag(`tag:"config"`)

	tag := ParseStructTag(rawTag)

	assert.True(t, len(tag["tag"]) == 1)
	assert.True(t, tag["tag"][0] == "config")
}