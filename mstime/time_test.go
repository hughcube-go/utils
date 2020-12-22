package mstime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Parse(t *testing.T) {
	a := assert.New(t)

	now := time.Now()

	date := now.Format(YYYYMMDDHHIISS)
	parseNow, err := Parse(date)
	a.Nil(err)
	a.Equal(now.Unix(), parseNow.Unix())
}

func Test_Format(t *testing.T) {
	a := assert.New(t)

	now := time.Now()

	date := Format(now.Local())
	parseNow, err := Parse(date)
	a.Nil(err)
	a.Equal(now.Unix(), parseNow.Unix())
}

func Test_FormatMst(t *testing.T) {
	a := assert.New(t)

	now := time.Now()

	date := FormatMst(now.Local())
	println(date)
	parseNow, err := ParseMst(date)
	a.Nil(err)
	a.Equal(now.Unix(), parseNow.Unix())
}
