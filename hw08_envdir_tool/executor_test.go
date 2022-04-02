package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		var env Environment
		cmd := []string{
			"test",
		}

		ComArgVar := RunCmd(cmd, env)
		require.Equal(t, 1, ComArgVar)
	})

	t.Run("0", func(t *testing.T) {
		var env Environment
		cmd := []string{
			"test",
			"-test",
		}

		ComArgVar := RunCmd(cmd, env)
		require.Equal(t, 0, ComArgVar)
	})

	t.Run("-1", func(t *testing.T) {
		var env Environment
		cmd := []string{
			"-test",
		}

		ComArgVar := RunCmd(cmd, env)
		require.Equal(t, -1, ComArgVar)
	})
}
