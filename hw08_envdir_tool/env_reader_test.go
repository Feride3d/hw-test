package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	envs, err := ReadDir("testdata/env/")
	require.NoError(t, err)
	expected := make(Environment)
	for _, kv := range [][2]string{
		{"BAR", "bar"},
		{"FOO", "   foo\nwith new line"},
		{"HELLO", `"hello"`},
	} {
		expected[kv[0]] = kv[1]
	}
	require.Equal(t, expected, envs)
}
