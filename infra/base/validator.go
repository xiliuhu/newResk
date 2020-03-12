package base

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	vtzh "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Translate() ut.Translator {
	return translator
}

func Init() {
	//创建validator对象
	validate := validator.New()
	//中文翻译器
	cn := zh.New()
	//通用翻译器
	uni := ut.New(cn, cn)
	//获取通用中文翻译器
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		//注册到验证器
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logrus.Info(err)
		}
	} else {
		logrus.Info("not find translator:zh")
	}
	/*	//验证
		err:=validate.Struct(user)
		//判断验证的结果
		if err !=nil{
			if _,ok:= err.(*validator.InvalidValidationError);ok{
				logrus.Info(err)
			}
			if errs,ok:=err.(validator.ValidationErrors);ok{
				for _,err:=range errs{
					logrus.Info(err.Translate(translator))
				}
			}
		}*/
}
