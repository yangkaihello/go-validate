package validate

import (
	"fmt"
	"strconv"
	"strings"
)

type Length struct {
	message string
	alias string
	name string
	rule string
	data interface{}
}

func (this *Length) SetAlias(s string) {
	this.alias = s
}

func (this *Length) GetAlias() string {
	return this.alias
}

func (this *Length) SetName(s string) {
	this.name = s
}

func (this *Length) GetName() string {
	return this.name
}

func (this *Length) SetRule(s string) {
	this.rule = s
}

func (this *Length) GetRule() string {
	return this.rule
}

func (this *Length) SetData(i interface{}) {
	this.data = i
}

func (this *Length) GetData() interface{} {
	return this.data
}

func (this *Length) SetErrorMessage(message string) {
	this.message = message
}

func (this *Length) GetErrorMessage() string {
	if this.message == "" {
		this.message = "%s超出%s验证范围(%s)"
	}

	var infoSlice = []interface{}{this.GetAlias(),this.GetRule(),this.GetData().(string)}
	infoSlice = infoSlice[0:strings.Count(this.message,"%s")]
	return fmt.Sprintf(this.message,infoSlice...)
}

func (this *Length) Verify() bool {
	var data = this.GetData()
	var length = strings.Split(this.GetRule(),",")

	if len(length) != 2 {
		this.SetErrorMessage("验证格式不正确，请使用1,1")
		return false
	}

	startLen,_ := strconv.Atoi(length[0])
	endLen,_ := strconv.Atoi(length[1])
	dataLen := len([]byte(data.(string)))

	if data == nil || startLen > dataLen || endLen < dataLen {
		return false
	}else{
		return true
	}
}