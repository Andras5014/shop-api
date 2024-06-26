package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"shop_api/user_web/global"
	"strings"
)

func InitTrans(local string) (err error) {
	// 修改gin自带validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用（fallback）的语言环境,后面是支持的语言环境
		uni := ut.New(zhT, enT, zhT)
		global.Trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}
		switch local {
		case "en":
			err = entranslations.RegisterDefaultTranslations(v, global.Trans)
			if err != nil {
				return err
			}
		case "zh":
			err = zhtranslations.RegisterDefaultTranslations(v, global.Trans)
			if err != nil {
				return err
			}
		default:
			err = entranslations.RegisterDefaultTranslations(v, global.Trans)
			if err != nil {
				return err
			}
		}
		return
	}
	return
}
