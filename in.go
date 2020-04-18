package validate

import (
	"fmt"
	"strings"
)

type In struct {
	message string
	alias   string
	name    string
	rule    string
	data    interface{}
}

func (this *In) SetAlias(s string) {
	this.alias = s
}

func (this *In) GetAlias() string {
	return this.alias
}

func (this *In) SetName(s string) {
	this.name = s
}

func (this *In) GetName() string {
	return this.name
}

func (this *In) SetRule(s string) {
	this.rule = s
}

func (this *In) GetRule() string {
	return this.rule
}

func (this *In) SetData(i interface{}) {
	this.data = i
}

func (this *In) GetData() interface{} {
	return this.data
}

func (this *In) SetErrorMessage(message string) {
	this.message = message
}

func (this *In) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s不在%s范围内"
	}

	var infoSlice = []interface{}{this.GetAlias(), this.GetRule(), this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message, "%s")]
	return fmt.Sprintf(this.message, infoSlice...)
}

func (this *In) Verify() bool {
	var data = this.GetData()
	var length = strings.Split(this.GetRule(),",")

	var exist = true
	for _,v := range length {
		if v == data {
			exist = false
		}
	}

	if data == nil || exist {
		return false
	}else{
		return true
	}
}