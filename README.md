# go-errors

A simple Go error utility to record the error stack trace, which supports customized formatting and is easy to migrate.

# How to

Replace import paths of default package errors by:

```
import "github.com/IguteChung/go-errors"
```

Replace all the occurences of `fmt.Errorf` by:

```
errors.Errorf("something wrong: %v", err)
```

# You are done, let's get the stacktrace

```
fmt.Println(errors.StackTrace(err))
```

```sh
main.bar
	/tmp/sandbox154253539/prog.go:18
main.foo
	/tmp/sandbox154253539/prog.go:14
main.main
	/tmp/sandbox154253539/prog.go:9
```
