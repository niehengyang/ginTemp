package utils

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"reflect"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (bool, interface{}) {
	err := c.Bind(form)
	if err != nil {
		return false, err
	}
	// 注册自定义校验函数
	RegisterCustomValidators()

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return false, err
	}
	if !check {

		markErrors := make([]string, 0)
		for _, err := range valid.Errors {
			markErrors = append(markErrors, err.Message)
		}
		return false, markErrors
	}

	return true, nil
}

// Bool 自定义布尔类型指针校验函数
func Bool(v *validation.Validation, obj interface{}, key string) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() || val.Elem().Kind() != reflect.Bool {
		_ = v.SetError(key, fmt.Sprintf("%s must be a boolean", key))
	}
}

// RegisterCustomValidators 注册自定义校验函数
func RegisterCustomValidators() {
	_ = validation.AddCustomFunc("Bool", Bool)
}
