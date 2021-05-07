package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type Data struct {
	Name       string `json:"name,omitempty" db:"name" validate:"required" label:"姓名"`
	Age        uint   `json:"age,omitempty" db:"age" validate:"gte=0,lte=100" label:"年龄"`
	Password   string `json:"password,omitempty" db:"password" validate:"required,alphanum,min=6" label:"密码"`
	RePassword string `json:"re_password,omitempty" db:"re_password"  validate:"eqfield=Password" label:"重密密码"`
}

var (
	validate = validator.New()
	d        = &Data{
		Name:       "",
		Age:        101,
		Password:   "a123456",
		RePassword: "a123456",
	}
)

// 自定义验证，返回 bool 决定了验证是否通过
//	fl validator.FieldLevel 包含了字段的所有信息以及相关的助手函数
func MyValidation(fl validator.FieldLevel) bool {
	if fl.Field().String() == "jack" {
		return true
	} else {
		return false
	}
}

// 自定义验证示例
func customValidationExample() {

	type tmp struct {
		Name       string `json:"name,omitempty" db:"name" validate:"required,myValidation"`
		Age        uint   `json:"age,omitempty" db:"age" validate:"gte=0,lte=100"`
		Password   string `json:"password,omitempty" db:"password" validate:"required,alphanum,min=6"`
		RePassword string `json:"re_password,omitempty" db:"re_password"  validate:"eqfield=Password"`
	}

	t := &tmp{
		Name:       "pony",
		Age:        101,
		Password:   "a123456",
		RePassword: "a123456",
	}

	// RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) 用于注册自定义验证
	if err := validate.RegisterValidation("myValidation", MyValidation); err != nil {
		log.Fatalln(err)
	}

	if err := validate.Struct(t); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

// 基础的验证示例
func example() {
	if err := validate.Struct(d); err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			// v.Field() 表示验证失败的字段
			// v.Value() 表示验证失败所传入的值
			// v.Tag() 表示导致验证失败的规则
			// v.Param() 表示导致验证失败的规则所定义的参数，例如 gt=1，若是实际的值小于 1 ，则返回 1
			fmt.Println(v.Value())
		}
		return
	}

	fmt.Println("success")
}

// 错误信息转中文；注：错误信息为中文，但字段依然保持不变。
func zhExample() {

	uni := ut.New(zh.New()) // 万能翻译器，保存所有的语言环境和翻译数据
	trans, _ := uni.GetTranslator("zh")

	// 验证器注册翻译器
	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalln(err)
	}

	err := validate.Struct(d)
	if err != nil {
		// 获取经过翻译的验证错误 map
		// err.(validator.ValidationErrors).Translate(trans)

		// 逐条翻译
		for _, v := range err.(validator.ValidationErrors) {
			fmt.Println(v.Translate(trans))
		}
	}
}

// 注册自定义字段名
func registerTagNameExample() {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalln(err)
	}

	// 注册 tag 作为错误信息的字段名，此处设置 tag 为 label
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	if err := validate.Struct(d); err != nil {
		fmt.Println(err.(validator.ValidationErrors).Translate(trans))
	}

}

func main() {

	// registerTagNameExample()
}
