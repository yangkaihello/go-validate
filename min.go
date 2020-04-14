package validate

import (
	"fmt"
	"strconv"
	"strings"
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

	var infoSlice = []interface{}{this.GetAlias(),this.GetRule(),this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message,"%s")]
	return fmt.Sprintf(this.message,infoSlice...)
}

func (this *Min) Verify() bool {
	var data = this.GetData()
	rule,_ := strconv.Atoi(this.GetRule())

	if data == nil || data.(int) < rule {
		return false
	}else{
		return true
	}
}