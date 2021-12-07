package myerrors

type MyError struct {
	errorCode int
	message   string
}

func (e MyError) Error() string {
	return e.message
}

var (
	myError1 = MyError{1, "my error 1"}
	myError2 = MyError{2, "my error 2"}
	//MyError3 = MyError{3, "my error 3"}
	//MyError4 = MyError{4, "my error 4"}
)
