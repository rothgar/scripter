package main

import "os"
import "syscall"
import "os/exec"

func main() {

	args := []string{os.Args[1]}

	binary, lookErr := exec.LookPath(os.Args[1])
	if lookErr != nil {
		panic(lookErr)
	}

	for i := 3; i <= len(os.Args); i++ {
		args = append(args, os.Args[i-1])
	}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
