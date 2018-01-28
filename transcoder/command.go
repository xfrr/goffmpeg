package transcoder

import (
	"fmt"
	"log"
	"os/exec"
)

type Worker struct {
	Command string
	Args    string
	Output  chan string
}

func (cmd *Worker) Run() {
	out, err := exec.Command(cmd.Command, cmd.Args).Output()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Output <- string(out)
}

func Collect(c chan string) {
	for {
		msg := <-c
		fmt.Printf("The command result is %s\n", msg)
	}
}