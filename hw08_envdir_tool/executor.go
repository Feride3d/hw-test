package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	for envVar, val := range env {
		if val.NeedRemove {
			os.Unsetenv(envVar)
			break
		}

		if _, ok := os.LookupEnv(envVar); ok {
			os.Unsetenv(envVar)
		}

		os.Setenv(envVar, val.Value)
	}

	_ = command.Run()
	return command.ProcessState.ExitCode()
}
