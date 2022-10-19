package vertify

import (
	"fmt"
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	zh_tw_translations "github.com/go-playground/validator/v10/translations/zh_tw"
	"net/http"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

type User struct {
	Username string `form:"user_name" validate:"required"`
	Tagline  string `form:"tag_line" validate:"required,lt=10"`
	Tagline2 string `form:"tag_line2" validate:"required,gt=1"`
}

func StartPage(c *gin.Context) {
	locale := c.DefaultQuery("locale", "zh")
	trans, _ := Uni.GetTranslator(locale)

	switch locale {
	case "zh":
		if err := zh_translations.RegisterDefaultTranslations(Validate, trans); err != nil {
			fmt.Println(err.Error())
		}
		break
	case "en":
		if err := en_translations.RegisterDefaultTranslations(Validate, trans); err != nil {
			fmt.Println(err.Error())
		}
		break
	case "zh_tw":
		if err := zh_tw_translations.RegisterDefaultTranslations(Validate, trans); err != nil {
			fmt.Println(err.Error())
		}
		break
	}

	if err := Validate.RegisterTranslation("reuqired", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	}); err != nil {
		fmt.Println("run error")
	}

	user := User{}
	if er := c.ShouldBind(&user); er != nil {
		fmt.Println(er.Error())
	}
	fmt.Println(user)
	err := Validate.Struct(user)

	if err != nil {
		e := err.(validator.ValidationErrors)
		var sliceErrs []string

		for _, fe := range e {
			sliceErrs = append(sliceErrs, fe.Translate(trans))
		}
		c.String(http.StatusOK, fmt.Sprintf("%#v", sliceErrs))
	}
	c.String(http.StatusOK, fmt.Sprintf("%#v", "user"))
}

func HandleError() {
	en := en2.New()
	zh := zh2.New()
	zhTw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zhTw)
	Validate = validator.New()

	r := gin.Default()
	r.GET("/8848", StartPage)
	r.POST("/sss", StartPage)
	if err := r.Run(":900"); err != nil {
		fmt.Println("wrong!")
	}
}
