package main

import (
	"log"
	"os"
	"strings"

	"github.com/Nekroze/runto/pkg/executor"
)

func main() {
	runner := executor.Runner{
		Command:   strings.Join(os.Args[1:], " "),
		Exec:      executor.ExecuteToStdout,
		Condition: executor.StopOnSuccess,
	}
	result := runner.Loop()
	log.Println("Done\n" + result.String())
	os.Exit(result.LastExitCode)
}
