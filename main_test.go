package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSource(t *testing.T) {
	fpath := "./test/foo/bar.txt"
	sourceBody := []byte("hello test")

	t.Run("create dir and file", func(t *testing.T) {
		err := createSource(fpath, sourceBody, 0600)
		assert.Nil(t, err)

		// check if file was written successful
		data, err := ioutil.ReadFile(fpath)
		assert.Nil(t, err)
		assert.Equal(t, sourceBody, data)
	})

	t.Run("file already exist", func(t *testing.T) {
		err := createSource(fpath, sourceBody, 0600)
		assert.Nil(t, err)
	})

	// teardown
	err := os.RemoveAll("./test")
	assert.Nil(t, err)
}
