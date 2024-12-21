package lib

import (
	"fmt"
	"runtime"
)

func Trace(depth int) string {
	trace := ""
	for i := 1; i <= depth; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		trace += fmt.Sprintf("%s:%d %s\n", file, line, fn.Name())
	}
	if trace == "" {
		return "unknown"
	}
	return trace
}

func ExampleFunc() {
	fmt.Println("exampleFunc called from", Trace(3))
}
