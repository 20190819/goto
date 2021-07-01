package validate

import (
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"reflect"
	"strings"
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
		if structField.Type.Kind().String() == "struct" {
			getMap(validateMap, validateNameMap, structField.Type)
		}

		fieldName := structField.Name
		if structField.Tag.Get("json") != "" {
			fieldName = structField.Tag.Get("json")
		}
		if structField.Tag.Get("validate") == "" {
			continue
		}
		validateMap[fieldName] = strings.Split(structField.Tag.Get("validate"), "||")
		if structField.Tag.Get("fieldName") == "" {
			continue
		}
		validateNameMap[fieldName] = structField.Tag.Get("fieldName")
	}
}

// 执行校验
func handleValidate(r *http.Request, validateMap map[string][]string, messages map[string][]string) (bool, string) {
	rules := govalidator.Options{
		Request:         r,
		Rules:           validateMap,
		Messages:        messages,
		RequiredDefault: false,
	}

	v := govalidator.New(rules)
	e := v.Validate()

	if len(e) == 0 {
		return true, ""
	}

	for _, err := range e {
		return false, err[0]
	}

	return true, ""
}
