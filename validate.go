package validate

import (
	"errors"
	_interface "github.com/yangkaihello/go-validate/interface"
	"reflect"
	"strings"
)

var TAG_NAME = "validate"

type firstError string

func (this *firstError) Get(err map[string]map[string]string) string {
	if *this != "" {
		stringSplit := strings.Split(string(*this), ":")[0:2]
		if d, ok := err[stringSplit[0]][stringSplit[1]]; ok {
			return d
		}
	}
	return ""
}

func (this *firstError) Set(name string, rule string) {
	if *this == "" {
		*this = firstError(name + ":" + rule)
	}
}

func (this *firstError) Del() {
	*this = ""
}

type Validate struct {
	firstErrorName firstError
	errorMessages  map[string]string
	errorNames     map[string]string
	errors         map[string]map[string]string
	rule           map[string]map[string]_interface.Rule
	ruleTemp       map[string]_interface.Rule
	single         *Validate
}

func (this *Validate) New() *Validate {
	this.single = this
	this.single.rule = map[string]map[string]_interface.Rule{}
	this.single.ruleTemp = map[string]_interface.Rule{}
	this.single.errors = map[string]map[string]string{}
	this.single.errorNames = map[string]string{}
	this.single.errorMessages = map[string]string{}

	this.single.SetRule("require", &Require{})
	this.single.SetRule("property", &Property{})
	this.single.SetRule("min", &Min{})
	this.single.SetRule("max", &Max{})
	this.single.SetRule("ip", &Ip{})
	this.single.SetRule("length", &Length{})
	this.single.SetRule("email", &Email{})
	this.single.SetRule("mobile", &Mobile{})
	this.single.SetRule("url", &Url{})
	return this
}

func (this *Validate) SetRule(key string, rule _interface.Rule) {
	this.single.ruleTemp[key] = rule
}

func (this *Validate) SetMessage(message map[string]string) {
	for key, value := range message {
		this.single.errorMessages[key] = value
	}
}

func (this *Validate) LoadDate(s interface{}, data map[string]interface{}) error {
	var typeOf = reflect.TypeOf(s)
	var valueOf = reflect.Indirect(reflect.ValueOf(s))

	if typeOf.Kind() != reflect.Ptr {
		return errors.New("类型错误不是struct指针")
	}
	var typeOfElem = typeOf.Elem()
	var dataValue interface{}
	var dataName string
	var aliasName string
	var errorTemp string
	var ok bool

	for i := 0; i < typeOfElem.NumField(); i++ {

		if valueOf.Field(i).Kind() != reflect.String {
			return errors.New("数据类型只能是string,如果需要验证数据类型请使用property")
		}

		dataName = strings.ToLower(typeOfElem.Field(i).Name)

		aliasName = dataName
		if valueOf.Field(i).String() != "" {
			aliasName = valueOf.Field(i).String()
		}

		tagString := typeOfElem.Field(i).Tag.Get(TAG_NAME)
		tagStrings := strings.Split(tagString, "|")

		if dataValue, ok = data[dataName]; !ok {
			dataValue = ""
		}

		for _, value := range tagStrings {
			rules := strings.Split(value, ":")
			ruleName := rules[0]

			if _, ok := this.single.rule[dataName]; !ok {
				this.single.rule[dataName] = map[string]_interface.Rule{}
			}

			if errorTemp, ok = this.single.errorMessages[dataName+":"+ruleName]; !ok {
				errorTemp = this.single.errorMessages[ruleName]
			}

			switch len(rules) {
			case 1:
				ruleInterface := reflect.New(reflect.TypeOf(this.single.ruleTemp[ruleName]).Elem()).Interface()
				rule := ruleInterface.(_interface.Rule)

				rule.SetErrorMessage(errorTemp)
				rule.SetAlias(aliasName)
				rule.SetName(dataName)
				rule.SetData(dataValue)
				this.single.rule[dataName][ruleName] = rule

			case 2:
				ruleInterface := reflect.New(reflect.TypeOf(this.single.ruleTemp[ruleName]).Elem()).Interface()
				rule := ruleInterface.(_interface.Rule)

				rule.SetErrorMessage(errorTemp)
				rule.SetAlias(aliasName)
				rule.SetName(dataName)
				rule.SetData(data[dataName])
				rule.SetRule(rules[1])
				this.single.rule[dataName][ruleName] = rule
			}

		}

	}

	return nil
}

func (this *Validate) Ok() bool {
	for name, value := range this.single.rule {
		for rule, single := range value {
			if single.Verify() == false {

				//记录所有错误
				if _, ok := this.single.errors[name]; !ok {
					this.single.errors[name] = map[string]string{}
				}
				this.single.errors[name][rule] = single.GetErrorMessage()

				//记录每个数据的首次错误
				if this.single.errorNames[name] == "" {
					this.single.errorNames[name] = single.GetErrorMessage()
				}

				//第一次错误的时候记录所属键值
				this.firstErrorName.Set(name, rule)

			}

		}
	}

	if len(this.single.errors) == 0 {
		return true
	} else {
		return false
	}

}

func (this *Validate) GetFirstError() string {
	return this.firstErrorName.Get(this.errors)
}

func (this *Validate) GetAllError() map[string]string {
	return this.single.errorNames
}

func (this *Validate) GetError(name string) map[string]string {
	var errors = map[string]string{}

	if d, ok := this.single.errors[name]; ok {
		errors = d
	}

	return errors
}
