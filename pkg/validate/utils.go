package validate

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

func getValidateMap(data interface{}) (map[string][]string, map[string]string) {
	typeOf := reflect.TypeOf(data)
	validateMap := make(map[string][]string)
	validateNameMap := make(map[string]string)
	getMap(validateMap, validateNameMap, typeOf)
	return validateMap, validateNameMap
}

func getMap(validateMap map[string][]string, validateNameMap map[string]string, typeOf reflect.Type) {
	for i := 0; i < typeOf.NumField(); i++ {
		structField := typeOf.Field(i)
		var fieldName string
		if structField.Tag.Get("json") != "" {
			fieldName = structField.Tag.Get("json")
		} else {
			fieldName = structField.Name
		}
		if structField.Type.Kind().String() == "struct" {
			getMap(validateMap, validateNameMap, structField.Type)
		} else {
			validateContent := structField.Tag.Get("validate")
			if validateContent != "" {
				if structField.Tag.Get("json") != "" {
					fieldName = structField.Tag.Get("json")
				}
				validateMap[fieldName] = strings.Split(validateContent, "||")
				if structField.Tag.Get("fieldName") != "" {
					validateNameMap[fieldName] = structField.Tag.Get("fieldName")
				}
			}
		}
	}
}

// 执行校验
func handleValidate(r *http.Request, validateMap map[string][]string, messages govalidator.MapData) (bool, string) {
	rules := govalidator.Options{
		Request:         r,
		Rules:           validateMap,
		Messages:        messages,
		RequiredDefault: true,
	}

	v := govalidator.New(rules)
	e := v.Validate()
	fmt.Println("eee>>>", e)
	if len(e) == 0 {
		return true, ""
	}

	for _, err := range e {
		return false, err[0]
	}

	return true, ""
}
