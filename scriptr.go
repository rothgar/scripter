package main

import "os"
import "syscall"
import "os/exec"

func main() {

	// set executable passed as first arg
	args := []string{os.Args[1]}

	// look for the executeble in the path and reletive to scriptr path
	binary, lookErr := exec.LookPath(os.Args[1])
	if lookErr != nil {
		panic(lookErr)
	}

	// add args if more were passed
	for i := 3; i <= len(os.Args); i++ {
		args = append(args, os.Args[i-1])
	}

	// load current environment
	env := os.Environ()

	// execute and look for errors
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
