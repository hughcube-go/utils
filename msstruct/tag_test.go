package msstruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tag_ParseTagWithSep(t *testing.T) {
	var err error

	a := assert.New(t)

	tag := ParseTagWithSep("primaryKey;is:;column:pk;sort:1;test:test:test:test", ";")

	a.True(tag.Has("primaryKey"))
	a.True(tag.Has("PRIMARYKEY"))
	a.True(tag.IsTrue("primaryKey"))

	a.True(tag.Has("is"))
	a.False(tag.IsTrue("is"))

	a.True(tag.Has("column"))
	a.False(tag.IsTrue("column"))
	a.Equal("pk", tag.Get("column"))

	a.True(tag.Has("sort"))
	a.True(tag.IsTrue("sort"))
	a.Equal("1", tag.Get("sort"))

	a.True(tag.Has("test"))
	a.False(tag.IsTrue("test"))
	a.Equal("test:test:test", tag.Get("test"))

	_, err = tag.GetInt("column")
	a.NotNil(err)
	sortInt, err := tag.GetInt("sort")
	a.Nil(err)
	a.Equal(sortInt, 1)
	a.IsType(sortInt, int(1))

	_, err = tag.GetInt64("column")
	a.NotNil(err)
	sortInt64, err := tag.GetInt64("sort")
	a.Nil(err)
	a.Equal(sortInt64, int64(1))
	a.IsType(sortInt64, int64(1))

	_, err = tag.GetFloat("column")
	a.NotNil(err)
	sortFloat, err := tag.GetFloat("sort")
	a.Nil(err)
	a.Equal(sortFloat, float32(1))
	a.IsType(sortFloat, float32(1))

	_, err = tag.GetFloat64("column")
	a.NotNil(err)
	sortFloat64, err := tag.GetFloat64("sort")
	a.Nil(err)
	a.Equal(sortFloat64, float64(1))
	a.IsType(sortFloat64, float64(1))
}
