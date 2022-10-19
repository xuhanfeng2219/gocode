package vertify

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"time"
)

type Person2 struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

func NameNotNullAndAdmin(
	v *validator.Validate,
	topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string) bool {

	if value, ok := field.Interface().(string); ok {
		return value != "" && !("ddj" == value)
	}

	return true
}

func ValidateNotNil() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("NotNullAndAdmin", func(fl validator.FieldLevel) bool {
			if value, ok := fl.Field().Interface().(string); ok {
				return value != "" && "dd" != value
			}
			return true
		}); err != nil {
			fmt.Println("run error")
		}
	}

	r.GET("/51hml", func(c *gin.Context) {
		var person Person2
		if err := c.ShouldBind(&person); err == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err:v%", err.Error())
		}
	})

	if err := r.Run(":90009"); err != nil {
		fmt.Println("name string get wrong!")
	}
}

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,BookableDate" time_format:"2016-10-01"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2016-10-02"`
}

func BookableDate(v *validator.Validate,
	topStruct reflect.Value,
	currentStructField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() > date.Unix() {
			return false
		}
	}
	return true
}

func GetBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "booking dates are valid!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
}

func HandleBooking() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterAlias("GetBookable", "BookableDate")
	}

	route.GET("/jdjjd", GetBookable)
	if err := route.Run(); err == nil {
		fmt.Println("is fine!")
	}
}
