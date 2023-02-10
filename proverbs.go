package commands

import "math/rand"

func RandomProverb() *Wisdom {
	i := rand.Intn(len(Wisdoms))
	return Wisdoms[i]
}

var proverbs = []*Wisdom{
	{"Don't communicate by sharing memory, share memory by communicating.", "Rob Pike"},
	{"Concurrency is not parallelism.", "Rob Pike"},
	{"Channels orchestrate; mutexes serialize.", "Rob Pike"},
	{"The bigger the interface, the weaker the abstraction.", "Rob Pike"},
	{"Make the zero value useful.", "Rob Pike"},
	{"interface{} says nothing.", "Rob Pike"},
	{"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.", "Rob Pike"},
	{"A little copying is better than a little dependency.", "Rob Pike"},
	{"Syscall must always be guarded with build tags.", "Rob Pike"},
	{"Cgo must always be guarded with build tags.", "Rob Pike"},
	{"Cgo is not Go.", "Rob Pike"},
	{"With the unsafe package there are no guarantees.", "Rob Pike"},
	{"Clear is better than clever.", "Rob Pike"},
	{"Reflection is never clear.", "Rob Pike"},
	{"Errors are values.", "Rob Pike"},
	{"Don't just check errors, handle them gracefully.", "Rob Pike"},
	{"Design the architecture, name the components, document the details.", "Rob Pike"},
	{"Documentation is for users.", "Rob Pike"},
	{"Don't panic.", "Rob Pike"},
}
