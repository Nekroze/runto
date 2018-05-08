package main

import (
	"os"
	"strings"

	"github.com/Nekroze/runto/pkg/executor"
)

func main() {
	os.Exit((&executor.Runner{
		Command:   strings.Join(os.Args[1:], " "),
		Exec:      executor.ExecuteToStdout,
		Condition: executor.StopOnFailure,
	}).LoopCLI())
}
