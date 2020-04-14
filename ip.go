package validate

import (
	"fmt"
	"net"
	"strings"
)

type Ip struct {
	message string
	alias   string
	name    string
	rule    string
	data    interface{}
}

func (this *Ip) SetAlias(s string) {
	this.alias = s
}

func (this *Ip) GetAlias() string {
	return this.alias
}

func (this *Ip) SetName(s string) {
	this.name = s
}

func (this *Ip) GetName() string {
	return this.name
}

func (this *Ip) SetRule(s string) {
	this.rule = s
}

func (this *Ip) GetRule() string {
	return this.rule
}

func (this *Ip) SetData(i interface{}) {
	this.data = i
}

func (this *Ip) GetData() interface{} {
	return this.data
}

func (this *Ip) SetErrorMessage(message string) {
	this.message = message
}

func (this *Ip) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s格式不正确"
	}

	var infoSlice = []interface{}{this.GetAlias(), this.GetRule(), this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message, "%s")]
	return fmt.Sprintf(this.message, infoSlice...)
}

func (this *Ip) Verify() bool {

	var data = this.GetData()
	var address = net.ParseIP(data.(string))
	if data == nil || address == nil {
		return false
	} else {
		return true
	}
}
