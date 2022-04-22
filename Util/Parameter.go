package util

type Parameter struct {
	Key      string
	Value    interface{}
	Required bool
}

func NewParameter(key string, value interface{}, required bool) Parameter {
	return Parameter{Key: key, Value: value, Required: required}
}
