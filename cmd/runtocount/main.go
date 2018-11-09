package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/Nekroze/runto/pkg/executor"
)

func main() {
	runs := 10
	var err error
	if c := os.Getenv("RUNTOCOUNT"); c != "" {
		runs, err = strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
	}

	os.Exit((&executor.Runner{
		Command:   strings.Join(os.Args[1:], " "),
		Exec:      executor.ExecuteToStdout,
		Condition: executor.StopAfterRuns(runs),
	}).LoopCLI())
}
