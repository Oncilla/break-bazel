# Sample causing internal compiler error

This is a minimal sample that causes an internal compiler error
when running tests in the `breaker` and the `breaker_test` package.

To get the internal compiler error, simply run `make`:
````shell
make  # Runs bazel test //breaker/...

ERROR: ~/go/src/github.com/Oncilla/break-bazel/breaker/BUILD.bazel:10:1: GoCompile breaker/linux_amd64_stripped/go_default_test%/github.com/Oncilla/break-ba
zel/breaker_test.a failed (Exit 1) builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix linux_amd64 -src breaker/breaker.go -src breaker/breaker_private_test.go -src breaker/breaker_test.go -arc ... (remaining 16 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox
compile: error running compiler: exit status 2
~/.cache/bazel/_bazel_roos/f7f55e86af501fefe0a0a3128549f885/sandbox/linux-sandbox/1/execroot/__main__/breaker/breaker_test.go:10:2: internal compiler error:
 conflicting package heights 12 and 0 for path "github.com/Oncilla/break-bazel/breaker"

Please file a bug report including a short program that triggers the error.
https://golang.org/issue/new
INFO: Elapsed time: 0.212s, Critical Path: 0.09s
INFO: 0 processes.
FAILED: Build did NOT complete successfully

FAILED: Build did NOT complete successfully
makefile:8: recipe for target 'test' failed
make: *** [test] Error 1
````

Running the tests from the regular go tool chain does not cause an error:
````shell
make gotest  # Runs go test ./breaker/...
````

When only using the package `breaker_test` for the tests,
the internal compiler error does not occur:
````shell
make fixtest
````


The mock file and build files are auto generated:
````shell
make mocks
make gazelle
````


