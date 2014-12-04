package main

import "os"
import "syscall"
import "os/exec"
import "fmt"

func main() {

	args := []string{os.Args[1]}

	fmt.Println(args)
	binary, lookErr := exec.LookPath(os.Args[1])
	if lookErr != nil {
		panic(lookErr)
	}

	if os.Args[2] != "" {
		args = append(args, os.Args[2])
		fmt.Println(args)
	}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
