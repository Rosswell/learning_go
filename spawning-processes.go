/**
examples for spawning new, non-Go processes
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	// simple command that takes no arguments or input, and just prints something to stdout. the exec.Command helper
	// creates an object to represent this external process
	dateCmd := exec.Command("date")

	// .Output is another helper that handles the common case of running a command, waiting for it to finish, and
	// collecting its output. If there were no errors, dateOut will hold bytes with teh date info
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// now piping data to the external process on its stdin and collecting the results from its stdout
	grepCmd := exec.Command("grep", "hello")

	// explicitly gradding input/output pipes, starting the process, writing some input to it, reading the resulting
	// output, and finally writing for the process to exit
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// above omitted error checks, but the usual if err != nil pattern could've been used for all of them. also, only
	// StdoutPipe results were collected, but StderrPipe could've been collected in the exact same way
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// when spawning commands, we need to provide an explicitly delineated commands and argument arrays, vs. bein able
	// to just pass in one command-line string. If you want to spawn a full command with a string, you can use bash's
	// -c option:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
