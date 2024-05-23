package transform

import "strings"

func StringTrimSpace(s interface{}) (interface{}, error) {
	trimSpace := strings.TrimSpace(s.(string))
	return trimSpace, nil
}
