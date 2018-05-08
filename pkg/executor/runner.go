package executor

import (
	"log"
	"os"
)

type Runner struct {
	Command   string
	Exec      Executor
	Condition RunnerStopCheck
	result    *RunnerResult
}

func (r *Runner) exec() {
	log.Printf("%s iteration %d\n", os.Args[0], r.result.Attempts)
	result := r.Exec(r.Command)
	r.result.Record(result)
}

func (r *Runner) Loop() *RunnerResult {
	r.result = &RunnerResult{
		Command:     r.Command,
		Performance: &TimingStatistics{},
	}
	for {
		r.exec()
		if r.Condition(r.result) {
			return r.result
		}
	}
}

// Executes Runner.Loop but prints output instead of reporting it and returns exit code.
func (r *Runner) LoopCLI() int {
	result := r.Loop()
	log.Println("Done\n" + result.String())
	return result.LastExitCode
}
