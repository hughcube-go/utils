package msslice

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Slice_InArray(t *testing.T) {
	a := assert.New(t)

	a.True(InArray(1, []int{1, 2, 4}))
	a.False(InArray("1", []int{1, 2, 4}))

	items := [][]int{}
	a.False(InArray([]int{1, 2, 4}, items))

	items = append(items, []int{2, 4})
	a.False(InArray([]int{1, 2, 4}, items))

	items = append(items, []int{1, 2, 4})
	a.True(InArray([]int{1, 2, 4}, items))

	a.False(InArray([]int{1, 2, 4}, 1))
}

func Test_Slice_Search(t *testing.T) {
	a := assert.New(t)

	a.Equal(0, Search(1, []int{1, 2, 4}))
	a.Equal(-1, Search("1", []int{1, 2, 4}))

	items := [][]int{}
	a.Equal(-1, Search([]int{1, 2, 4}, items))

	items = append(items, []int{2, 4})
	a.Equal(-1, Search([]int{1, 2, 4}, items))

	items = append(items, []int{1, 2, 4})
	a.Equal(1, Search([]int{1, 2, 4}, items))
}

func Test_Slice_In(t *testing.T) {
	a := assert.New(t)

	a.True(In(1, 1, 2, 4))
	a.False(In("1", 1, 2, 4))

	a.False(In([]int{1, 2, 4}))
	a.False(In([]int{1, 2, 4}, []int{1, 2, 1}))
	a.True(In([]int{1, 2, 4}, []int{1, 2, 4}, []int{1, 2, 1}))
	a.False(In([]int{1, 2, 4}, 1))
}

func Test_Slice_GetElemType(t *testing.T) {
	var err error
	var typ reflect.Type

	a := assert.New(t)

	emptyInterfaceSlice := []interface{}{}
	typ, err = GetElemType(emptyInterfaceSlice, false)
	a.Nil(err)
	a.True(typ.Kind() == reflect.Interface)

	interfaceSlice := []interface{}{1}
	typ, err = GetElemType(interfaceSlice, false)
	a.Nil(err)
	a.True(typ.Kind() == reflect.Interface)

	intSlice := []int{1}
	typ, err = GetElemType(intSlice, false)
	a.Nil(err)
	a.True(typ.Kind() == reflect.Int)

	ptrIntSlice := []*int{}
	typ, err = GetElemType(ptrIntSlice, false)
	a.Nil(err)
	a.True(typ.Kind() == reflect.Ptr)

	typ, err = GetElemType(ptrIntSlice, true)
	a.Nil(err)
	a.True(typ.Kind() == reflect.Int)

	_, err = GetElemType(1, true)
	a.NotNil(err)
}

func Test_Slice_MakeSameTypeValue(t *testing.T) {
	a := assert.New(t)

	var newSliceValue reflect.Value
	var err error

	intSlice := []int{1}
	newSliceValue, err = MakeSameTypeValue(intSlice, 0, 0)
	a.Nil(err)
	a.IsType(intSlice, newSliceValue.Interface())

	structSlice := []struct{}{}
	newSliceValue, err = MakeSameTypeValue(structSlice, 0, 0)
	a.Nil(err)
	a.IsType(structSlice, newSliceValue.Interface())

	_, err = MakeSameTypeValue(1, 0, 0)
	a.NotNil(err)
}
