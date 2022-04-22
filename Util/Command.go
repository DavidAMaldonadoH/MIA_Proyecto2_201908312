package util

type Command interface {
	Execute() interface{}
}
