package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)

	for _, f := range files {
		filePath := path.Join(dir, f.Name())

		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			break
		}

		line = strings.TrimRight(line, " \t\n")
		line = string(bytes.ReplaceAll([]byte(line), []byte("\x00"), []byte("\n")))
		envs[f.Name()] = EnvValue{
			Value:      line,
			NeedRemove: false,
		}

		file.Close()
	}

	return envs, nil
}
