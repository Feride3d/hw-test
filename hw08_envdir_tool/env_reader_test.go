package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		envs, _ := ReadDir("testdata/env/")
		expected := Environment{
			"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY": EnvValue{Value: "", NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: false},
		}
		require.Equal(t, expected, envs)
	})

	t.Run("no directory", func(t *testing.T) {
		_, err := ReadDir("test.sh")
		require.NotEqual(t, nil, err)
	})
}
