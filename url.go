package validate

import (
	"fmt"
	"net/url"
	"strings"
)

type Url struct {
	message string
	alias   string
	name    string
	rule    string
	data    interface{}
}

func (this *Url) SetAlias(s string) {
	this.alias = s
}

func (this *Url) GetAlias() string {
	return this.alias
}

func (this *Url) SetName(s string) {
	this.name = s
}

func (this *Url) GetName() string {
	return this.name
}

func (this *Url) SetRule(s string) {
	this.rule = s
}

func (this *Url) GetRule() string {
	return this.rule
}

func (this *Url) SetData(i interface{}) {
	this.data = i
}

func (this *Url) GetData() interface{} {
	return this.data
}

func (this *Url) SetErrorMessage(message string) {
	this.message = message
}

func (this *Url) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s格式不正确"
	}

	var infoSlice = []interface{}{this.GetAlias(), this.GetRule(), this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message, "%s")]
	return fmt.Sprintf(this.message, infoSlice...)
}

func (this *Url) Verify() bool {

	var data = this.GetData()
	var _,err = url.Parse(data.(string))
	if data == nil || err != nil {
		return false
	} else {
		return true
	}
}
