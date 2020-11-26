package msfile

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test_Exists(t *testing.T) {
	a := assert.New(t)

	tempDir := os.TempDir()

	a.True(Exists(tempDir))
	a.False(Exists(tempDir + strconv.FormatInt(time.Now().UnixNano(), 10)))

	file, err := ioutil.TempFile("", "")
	a.Nil(err)

	a.True(Exists(file.Name()))
	a.False(Exists(file.Name() + strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func Test_IsDir(t *testing.T) {
	a := assert.New(t)

	tempDir := os.TempDir()

	a.True(IsDir(tempDir))
	a.False(IsDir(tempDir + strconv.FormatInt(time.Now().UnixNano(), 10)))

	file, err := ioutil.TempFile("", "")
	a.Nil(err)

	a.False(IsDir(file.Name()))
	a.False(IsDir(file.Name() + strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func Test_IsFile(t *testing.T) {
	a := assert.New(t)

	tempDir := os.TempDir()

	a.False(IsFile(tempDir))
	a.False(IsFile(tempDir + strconv.FormatInt(time.Now().UnixNano(), 10)))

	file, err := ioutil.TempFile("", "")
	a.Nil(err)

	a.True(IsFile(file.Name()))
	a.False(IsFile(file.Name() + strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func Test_Size(t *testing.T) {
	a := assert.New(t)

	tempDir := os.TempDir()

	a.GreaterOrEqual(Size(tempDir), int64(0))
	a.Equal(Size(tempDir+strconv.FormatInt(time.Now().UnixNano(), 10)), int64(0))

	file, err := ioutil.TempFile("", "")

	a.Nil(err)

	a.Equal(Size(file.Name()), int64(0))
	a.Equal(Size(file.Name()+strconv.FormatInt(time.Now().UnixNano(), 10)), int64(0))

	size, err := file.Write([]byte("你好"))
	a.Nil(err)

	a.Equal(Size(file.Name()), int64(size))

	if file != nil {
		_ = file.Close()
		_ = os.RemoveAll(file.Name())
		a.False(Exists(file.Name()))
	}
}

func Test_Mkdir(t *testing.T) {
	a := assert.New(t)

	tempDir := os.TempDir()

	dirs := []string{
		tempDir + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "/111/",
		tempDir + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "/1/1/1/",
		tempDir + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "/222",
		tempDir + "/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "/2/2/2/",
	}

	for _, v := range dirs {
		err := Mkdir(v, 0777)
		a.Nil(err)
		a.True(Exists(v))
		_ = os.RemoveAll(v)
		a.False(Exists(v))
	}
}
