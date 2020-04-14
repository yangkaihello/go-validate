package validate

import (
	"fmt"
	"regexp"
	"strings"
)

type Email struct {
	message string
	alias   string
	name    string
	rule    string
	data    interface{}
}

func (this *Email) SetAlias(s string) {
	this.alias = s
}

func (this *Email) GetAlias() string {
	return this.alias
}

func (this *Email) SetName(s string) {
	this.name = s
}

func (this *Email) GetName() string {
	return this.name
}

func (this *Email) SetRule(s string) {
	this.rule = s
}

func (this *Email) GetRule() string {
	return this.rule
}

func (this *Email) SetData(i interface{}) {
	this.data = i
}

func (this *Email) GetData() interface{} {
	return this.data
}

func (this *Email) SetErrorMessage(message string) {
	this.message = message
}

func (this *Email) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s格式不正确"
	}

	var infoSlice = []interface{}{this.GetAlias(), this.GetRule(), this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message, "%s")]
	return fmt.Sprintf(this.message, infoSlice...)
}

func (this *Email) Verify() bool {
	var data = this.GetData()

	pattern := `^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`
	reg := regexp.MustCompile(pattern)

	if data == nil || !reg.MatchString(data.(string)) {
		return false
	} else {
		return true
	}
}
