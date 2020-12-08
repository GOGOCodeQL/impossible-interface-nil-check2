package main

//MyError is My Error :)
type MyError struct {
	Op   string
	Path string
	Err  error
}

func (E MyError) Error() string {
	return E.Op + "bye!!!"
}

func returnsError() *MyError {
	var p *MyError = nil
	return p // Will always return a non-nil error.
}

func main() {
	var e error = returnsError()
	if e != nil {
		println("Error!!")
	} else {
		println("404 No Error") //unreachable
	}
}
