package validate

import (
	"fmt"
	"strings"
)

type Mobile struct {
	message string
	alias   string
	name    string
	rule    string
	data    interface{}
}

func (this *Mobile) SetAlias(s string) {
	this.alias = s
}

func (this *Mobile) GetAlias() string {
	return this.alias
}

func (this *Mobile) SetName(s string) {
	this.name = s
}

func (this *Mobile) GetName() string {
	return this.name
}

func (this *Mobile) SetRule(s string) {
	this.rule = s
}

func (this *Mobile) GetRule() string {
	return this.rule
}

func (this *Mobile) SetData(i interface{}) {
	this.data = i
}

func (this *Mobile) GetData() interface{} {
	return this.data
}

func (this *Mobile) SetErrorMessage(message string) {
	this.message = message
}

func (this *Mobile) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s超出%s"
	}

	var infoSlice = []interface{}{this.GetAlias(), this.GetRule(), this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message, "%s")]
	return fmt.Sprintf(this.message, infoSlice...)
}

func (this *Mobile) Verify() bool {
	var data = this.GetData()

	if data == nil || len([]byte(data.(string))) != 11 {
		return false
	} else {
		return true
	}
}
