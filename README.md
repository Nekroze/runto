# runto

Utility to rerun command until a desired outcome. Currently this provides the following binaries that can be executed on the command line:

- `runtodeath` keep iterating until the command fails
- `runtolife` keep iterating until the command succeeds
- `runtocancel` keeps iterating until you press CTRL+C
- `runtocount` keeps iterating until it has reached `$RUNTOCOUNT` which defaults to `10`

# Install

You can install this the same as most go tools, via the `go get` command, which will install all binaries offered by this repository.

```bash
 $ go get -u github.com/Nekroze/runto/cmd/...
```

# Usage

If for example there is a command that checks that a service is online and you want to know when it is up, you may want to use `runtolife` command as a prefix to the check command to rerun it until the command has an exit code of 1. Another example use case is a flakey test when writing software, sometimes there is a rare edge case that may require rerunning a test many times before you might see the failure, `runtodeath` can help by runnning the test command for you until it fails.

# Output

This repository has a little script that randomly fails sometimes that we can use to play with the commands and see their output. One of the main advantages of using `runto` commands instead of something like a bash for loop or until builtin is that you get a report at the end of the execution about the run and the time the executions took:

```
2018/05/09 04:41:39 runtodeath iteration 0
Random success or failure
2018/05/09 04:41:40 runtodeath iteration 1
Random success or failure
2018/05/09 04:41:43 Done
# Command
Command:   ./.randexit.sh
Attempts:  2
Last Exit: 1
# Performance
Fastest: 1.00465452s
Slowest: 3.004678222s
Average: 2.004666371s
Total:   4.009332742s
```
