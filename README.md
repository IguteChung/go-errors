# go-errors

A simple Go error utility to record the error stack trace, which supports customized formatting and is easy to migrate.

# How to

Replace import paths of package errors:

```
import "github.com/IguteChung/go-errors"
```

Replace all the occurences of `fmt.Errorf`

```
errors.Errorf("something wrong: %v", err)
```

# You are done, let's get the stacktrace

Cast an error to `ErrorTracer` to get the stacktrace.

```
if tracer, ok := err.(errors.ErrorTracer); ok {
    fmt.Println(tracer.Stack().Format(errors.GoLikeFormatter))
}
```

```sh
main.bar
	/tmp/sandbox154253539/prog.go:18
main.foo
	/tmp/sandbox154253539/prog.go:14
main.main
	/tmp/sandbox154253539/prog.go:9
```
