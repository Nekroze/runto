package executor

import (
	"os"
	"os/signal"
)

var sigInterrupt bool

func init() {
	go func() {
		sigs := make(chan os.Signal, 2)
		signal.Notify(sigs, os.Interrupt)
		for sig := range sigs {
			switch sig {
			case os.Interrupt:
				sigInterrupt = true
			}
		}
	}()
}

// Takes a runner result and returns if the runner should stop or not.
type RunnerStopCheck func(*RunnerResult) bool

// Returns true (stop the runner) when the last command returned non 0 exit code.
func StopOnFailure(result *RunnerResult) bool {
	return result.LastExitCode != 0
}

func StopOnSuccess(result *RunnerResult) bool {
	return result.LastExitCode == 0
}

func StopAfterCancel(_ *RunnerResult) bool {
	return sigInterrupt
}

var runCounter int

func StopAfterRuns(runs int) RunnerStopCheck {
	return func(_ *RunnerResult) bool {
		runCounter += 1
		return runCounter >= runs
	}
}
