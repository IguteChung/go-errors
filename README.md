# go-errors

A simple Go error utility to record the error stack trace, which supports customized formatting and is easy to migrate.

# How to

Replace import paths of `import "errors"` by:

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

Output:

```
/tmp/sandbox647637562/prog.go:24
/tmp/sandbox647637562/prog.go:19
/tmp/sandbox647637562/prog.go:14
```

# More customization

Apply custom format for stack trace:

```
func init() {
	errors.ApplyFormatter("foo\n\tfile.go:152\n")
}
```

Output:

```
main.bar
	/tmp/sandbox639852497/prog.go:24
main.foo
	/tmp/sandbox639852497/prog.go:19
main.main
	/tmp/sandbox639852497/prog.go:14
```
