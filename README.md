# go-validate
* go 的数据验证器

### 简易使用

```
//验证器实例化
var vali = new(validate.Validate).New()
var err error
//需要验证的数据
var data = map[string]interface{}{
"name":"yangkai"
"age":18
}   

//验证的结构体模版
var valiStruct = &struct {
    Name string `validate:"require|property:string"`
    age string `validate:"require"|min:17`
}{Name:"名称"}

//验证器数据载入
err = vali.LoadDate(valiStruct, data)
//数据结构验证错误，不会继续往下验证数据
if err != nil {
    log.Println(err+"loadData 过程中存在的异常，比如你提供的验证结构体模版数据类型不正确")
}

//验证数据格式是否正确
if vali.Ok() == false {
    //打印所有异常的数据
    log.Println(vali.GetAllError())
}
```

### 提供验证

```
require         \\必填
property:int    \\数据类型 int,string,fooler
min:0           \\最小整形
max:99          \\最大整形
length:10,10    \\字符串长度验证
ip              \\ipv4数据类型验证
email           \\邮箱验证
mobile          \\手机号验证
url             \\url验证
in:git,svn      \\允许的数据
```
