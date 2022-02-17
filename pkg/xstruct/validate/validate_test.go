package validate

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestValidate(t *testing.T) {
	type Score struct {
		Subject string `validate:"required"`
		Score   int    `validate:"gte=0,lte=100"`
	}

	type User struct {
		Name  string   `validate:"required,gt=3,lt=10"`
		Age   uint8    `validate:"gte=0,lte=150"`
		Email string   `validate:"required,email"`
		Color string   `validate:"iscolor"`
		Score []*Score `validate:"required,dive,required"`
	}

	Convey("Validate error", t, func() {
		user := &User{
			Name:  "k",
			Age:   230,
			Email: "abc",
			Color: "abc",
			Score: nil,
		}
		err := Get().Struct(user)
		errs, ok := err.(ValidationErrors)
		fmt.Println(reflect.TypeOf(err))
		So(ok, ShouldBeTrue)
		for _, err := range errs {
			t.Log(err)
		}
		So(len(errs), ShouldEqual, reflect.ValueOf(*user).NumField())
	})

	Convey("Validate ok", t, func() {
		user := &User{
			Name:  "kkkk",
			Age:   100,
			Email: "abc@abc.com",
			Color: "#000000",
			Score: []*Score{
				{Subject: "math", Score: 100},
			},
		}
		err := Get().Struct(user)
		So(err, ShouldBeNil)
	})

	Convey("Func", t, func() {
		Convey("Custom Validation Functions", func() {
			var customFunc = func(fl validator.FieldLevel) bool {
				if fl.Field().String() != "TT" {
					return false
				}
				return true
			}
			type customStruct struct {
				Name string `validate:"customTag"`
			}
			struct1 := customStruct{"TT"}
			struct2 := customStruct{"TT_NOT"}
			validate.RegisterValidation("customTag", customFunc)
			err := validate.Struct(struct1)
			So(err, ShouldBeNil)
			err = validate.Struct(struct2)
			So(err, ShouldNotBeNil)
		})
	})
}
