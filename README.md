# runto

Utility to rerun command until a desired outcome. Currently this provides the following binaries that can be executed on the command line:

- `runtodeath`
- `runtolife`

# Install

You can install this the same as most go tools, via the `go get` command, which will install all binaries offered by this repository.

```bash
 $ go get -u github.com/Nekroze/runto/cmd/...
```

# Usage

If for example there is a command that checks that a service is online and you want to know when it is up, you may want to use `runtolife` command as a prefix to the check command to rerun it until the command has an exit code of 1. Another example use case is a flakey test when writing software, sometimes there is a rare edge case that may require rerunning a test many times before you might see the failure, `runtodeath` can help by runnning the test command for you until it fails.
