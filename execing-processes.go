/**
rather than spawning an external process that is accessible to the running Go process, this exemplifies completely
replacing the current Go process with another (maybe non-Go) one. Uses Go's implementation of the exec function
*/
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// going to use ls as an example. Go requires an absolute path to the binary to execute, so using exec.LookPath
	// to find it (probably /bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec requires arguments in slice form (as opposed to one big string). This gives ls a few common arguments (note
	// first argument is the program name
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environ variables to use. loading them here
	env := os.Environ()

	// this is the actual ls call. if this call is successful, the execution of the process will end here and be
	// replaced by the /bin/ls -a -l -h process. if there is an error, we'll get a return value
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
