package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("args are not provided: <env_dir> <command ...>")
	}
	dir, command := os.Args[1], os.Args[2:]
	envs, err := ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	exitCode := RunCmd(command, envs)
	os.Exit(exitCode)

	/* 	envs, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	returnCode := RunCmd(os.Args[2:], envs)
	if returnCode != 0 {
		log.Printf("command exited with code %d\n", returnCode)
		log.Fatal()
	} */
}
