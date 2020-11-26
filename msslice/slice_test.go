package msslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InArray(t *testing.T) {
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

func Test_Search(t *testing.T) {
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
