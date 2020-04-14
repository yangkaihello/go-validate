package validate

import (
	"fmt"
	"strconv"
	"strings"
)

type Max struct {
	message string
	alias string
	name string
	rule string
	data interface{}
}

func (this *Max) SetAlias(s string) {
	this.alias = s
}

func (this *Max) GetAlias() string {
	return this.alias
}

func (this *Max) SetName(s string) {
	this.name = s
}

func (this *Max) GetName() string {
	return this.name
}

func (this *Max) SetRule(s string) {
	this.rule = s
}

func (this *Max) GetRule() string {
	return this.rule
}

func (this *Max) SetData(i interface{}) {
	this.data = i
}

func (this *Max) GetData() interface{} {
	return this.data
}

func (this *Max) SetErrorMessage(message string) {
	this.message = message
}

func (this *Max) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s超出%s"
	}

	var infoSlice = []interface{}{this.GetAlias(),this.GetRule(),this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message,"%s")]
	return fmt.Sprintf(this.message,infoSlice...)
}

func (this *Max) Verify() bool {
	var data = this.GetData()
	rule,_ := strconv.Atoi(this.GetRule())

	if data == nil || data.(int) > rule {
		return false
	}else{
		return true
	}
}