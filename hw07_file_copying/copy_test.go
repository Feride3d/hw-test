package main

import (
	"os"
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
		err := Copy("/dev/urandom", "/tmp", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})
	t.Run("file doesn't exist", func(t *testing.T) {
		err := Copy("input.txt", "/tmp", 0, 0)
		require.NotNil(t, err)
	})
	t.Run("copy directory", func(t *testing.T) {
		err := Copy("/tmp", "input.txt", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})
}
