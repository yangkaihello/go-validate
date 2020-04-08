package _interface


type Rule interface {
	SetAlias(s string)
	GetAlias() string
	SetName(s string)
	GetName() string
	SetRule(s string)
	GetRule() string
	SetData(i interface{})
	GetData() interface{}
	SetErrorMessage(s string)
	GetErrorMessage() string
	Verify() bool
}

