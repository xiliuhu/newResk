package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	vtzh "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       string `validate:"gte=0,lte=130""`
	Email     string `validate:"required,email"`
}

func main() {
	user := &User{
		FirstName: "firstname",
		LastName:  "lastname",
		Age:       "137",
		Email:     "zt@mail.com",
	}
	//创建validator对象
	validate := validator.New()
	//中文翻译器
	cn := zh.New()
	//通用翻译器
	uni := ut.New(cn, cn)
	//获取通用中文翻译器
	translator, found := uni.GetTranslator("zh")
	if found {
		//注册到验证器
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("not find")
	}
	//验证
	err := validate.Struct(user)
	//判断验证的结果
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errs {
				fmt.Println(err.Translate(translator))
			}
		}
	}

}
