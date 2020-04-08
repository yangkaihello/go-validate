package validate

import (
	"fmt"
	"strconv"
)

type Min struct {
	message string
	alias string
	name string
	rule string
	data interface{}
}

func (this *Min) SetAlias(s string) {
	this.alias = s
}

func (this *Min) GetAlias() string {
	return this.alias
}

func (this *Min) SetName(s string) {
	this.name = s
}

func (this *Min) GetName() string {
	return this.name
}

func (this *Min) SetRule(s string) {
	this.rule = s
}

func (this *Min) GetRule() string {
	return this.rule
}

func (this *Min) SetData(i interface{}) {
	this.data = i
}

func (this *Min) GetData() interface{} {
	return this.data
}

func (this *Min) SetErrorMessage(message string) {
	this.message = message
}

func (this *Min) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s小于%s"
	}
	return fmt.Sprintf(this.message,this.GetAlias(),this.GetRule())
}

func (this *Min) Verify() bool {
	var data = this.GetData()
	rule,_ := strconv.Atoi(this.GetRule())

	if data.(int) < rule {
		return false
	}else{
		return true
	}
}