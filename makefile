.PHONY: all clean

GAZELLE_MODE?=fix

all:  test

test: gazelle
	bazel test //breaker/...

gotest:
	go test ./breaker/...

fixtest:
	cp breaker/breaker_private_test.fix breaker/breaker_private_test.go
	bazel test //breaker/...
	git checkout -- breaker

clean:
	rm -rf bazel-*

mocks:
	mockgen -source=breaker/breaker.go -destination breaker/mock_breaker/smasher.go

gazelle:
	bazel run //:gazelle -- update -mode=$(GAZELLE_MODE) -index=false -external=external  ./breaker
