package main

import "fmt"

type ErrorCode int

const (
	ErrNotFound ErrorCode = iota
	ErrPermissionDenied
	ErrTimeout
	ErrInternal
	ErrXXX
)

func (e ErrorCode) String() string {
	switch e {
	case ErrNotFound:
		return "ErrNotFound"
	case ErrPermissionDenied:
		return "ErrPermissionDenied"
	case ErrTimeout:
		return "ErrTimeout"

	case ErrInternal:
		return "ErrInternal"

	default:
		return "Unknown Error"

	}

}
func main() {
	errCode := ErrXXX
	fmt.Println(errCode)
	errCode = ErrNotFound
	fmt.Println(errCode)

}
