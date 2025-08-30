package bar 

func Bar() {
	bar()
}

//go:noinline
func bar() {
	panic("from bar")
}
