package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("copy file to file", func(t *testing.T) {
		f, _ := os.Create("input.txt")
		defer os.Remove("input.txt")
		_, _ = f.WriteString("test")

		err := Copy("input.txt", "input.txt", 0, 0)
		require.NotNil(t, err)
	})
	t.Run("copy /dev/urandom", func(t *testing.T) {
		err := Copy("/tmp", "/dev/urandom", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})
	t.Run("file doesn't exist", func(t *testing.T) {
		err := Copy("/tmp", "input.txt", 0, 0)
		require.NotNil(t, err)
	})
	t.Run("copy directory", func(t *testing.T) {
		err := Copy("/tmp", "input.txt", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})
}

func TestCopyTestdata(t *testing.T) {
	const dataPath = "testdata"
	srcPath := filepath.Join(dataPath, "input.txt")
	for _, ts := range []struct {
		offset int64
		limit  int64
	}{
		{offset: 0, limit: 0},
		{offset: 0, limit: 10},
		{offset: 0, limit: 1000},
		{offset: 0, limit: 10000},
		{offset: 100, limit: 1000},
		{offset: 6000, limit: 1000},
	} {
		t.Run(fmt.Sprintf("%v_%v", offset, limit), func(t *testing.T) {
			dst, err := ioutil.TempFile(os.TempDir(), "test")
			require.NoError(t, err)
			defer dst.Close()
			err = Copy(srcPath, dst.Name(), ts.offset, ts.limit)
			require.NoError(t, err)
			expPath := filepath.Join(dataPath, fmt.Sprintf("out_offset%d_limit%d.txt", ts.offset, ts.limit))
			fmt.Println(expPath)
			expContent, err := ioutil.ReadFile(expPath)
			require.NoError(t, err)
			dstContent, err := ioutil.ReadFile(dst.Name())
			require.NoError(t, err)
			x := bytes.Compare
			require.Zero(t, x(expContent, dstContent))
		})
	}
}
