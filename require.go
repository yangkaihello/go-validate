package validate

import (
	"fmt"
)

type Require struct {
	message string
	alias string
	name string
	rule string
	data interface{}
}

func (this *Require) SetAlias(s string) {
	this.alias = s
}

func (this *Require) GetAlias() string {
	return this.alias
}

func (this *Require) SetName(s string) {
	this.name = s
}

func (this *Require) GetName() string {
	return this.name
}

func (this *Require) SetRule(s string) {
	this.rule = s
}

func (this *Require) GetRule() string {
	return this.rule
}

func (this *Require) SetData(i interface{}) {
	this.data = i
}

func (this *Require) GetData() interface{} {
	return this.data
}

func (this *Require) SetErrorMessage(message string) {
	this.message = message
}

func (this *Require) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s必填"
	}
	return fmt.Sprintf(this.message,this.GetAlias())
}

func (this *Require) Verify() bool {
	if this.GetData() == nil || this.GetData() == "" {
		return false
	}else{
		return true
	}
}