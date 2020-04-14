package validate

import (
	"fmt"
	"reflect"
	"strings"
)

type Property struct {
	message string
	alias   string
	name    string
	rule    string
	data    interface{}
}

func (this *Property) SetAlias(s string) {
	this.alias = s
}

func (this *Property) GetAlias() string {
	return this.alias
}

func (this *Property) SetName(s string) {
	this.name = s
}

func (this *Property) GetName() string {
	return this.name
}

func (this *Property) SetRule(s string) {
	this.rule = s
}

func (this *Property) GetRule() string {
	return this.rule
}

func (this *Property) SetData(i interface{}) {
	this.data = i
}

func (this *Property) GetData() interface{} {
	return this.data
}

func (this *Property) SetErrorMessage(message string) {
	this.message = message
}

func (this *Property) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s数据类型不正确"
	}

	var infoSlice = []interface{}{this.GetAlias(), this.GetRule(), this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message, "%s")]
	return fmt.Sprintf(this.message, infoSlice...)
}

func (this *Property) Verify() bool {
	valueOf := reflect.Indirect(reflect.ValueOf(this.GetData()))
	switch this.GetRule() {
	case "string":
		if valueOf.Kind() == reflect.String {
			return true
		}
	case "int":
		if valueOf.Kind() == reflect.Int ||
			valueOf.Kind() == reflect.Int8 ||
			valueOf.Kind() == reflect.Int16 ||
			valueOf.Kind() == reflect.Int32 ||
			valueOf.Kind() == reflect.Int64 {
			return true
		}
	case "float":
		if valueOf.Kind() == reflect.Float32 ||
			valueOf.Kind() == reflect.Float64 {
			return true
		}
	}
	return false
}
