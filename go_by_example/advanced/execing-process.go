package main

import (
	"os"
	"os/exec"
	"syscall"
)

// Sometimes we just want to completely replace the current go process with another one.
// To do this we'll use go's implementation of the classic `exe` function.

func main() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
