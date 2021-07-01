package validate

import (
	"net/http"
)

func MapValidate(r *http.Request, data map[string][]string, validateMsg ...map[string]string) (bool, string) {
	var messages map[string][]string
	if len(validateMsg) > 0 {
		messages = getChMessage(data, validateMsg[0])
	} else {
		messages = getChMessage(data, nil)
	}
	return handleValidate(r, data, messages)
}

func StructValidate(r *http.Request, data interface{}, language string) (bool, string) {
	validateMap, validateNameMap := getValidateMap(data)
	var messages map[string][]string
	if language == "zh" {
		messages = getChMessage(validateMap, validateNameMap)
	}
	return handleValidate(r, validateMap, messages)
}

func MapDataForStruct(data interface{}) (map[string][]string, map[string]string) {
	return getValidateMap(data)
}
