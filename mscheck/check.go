package mscheck

import (
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func IsTrue(val interface{}) bool {
	if v, ok := val.(bool); ok {
		return v
	}

	if v, ok := val.(string); ok {
		v = strings.ToLower(v)
		if "on" == v || "yes" == v || "true" == v || "1" == v {
			return true
		}
	}

	return false
}

func IsMatchRegexp(val string, pattern string) bool {
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(val)
}

func IsEmail(val string) bool {
	// @see http://www.regular-expressions.info/email.html
	pattern := "^[a-zA-Z0-9!#$%&\\'*+\\\\/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&\\'*+\\\\/=?^_`{|}~-]+)*@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?\\.)+[a-zA-Z0-9](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$"

	return IsMatchRegexp(val, pattern)
}

func IsUrl(val string) bool {
	_, err := url.ParseRequestURI(val)
	return err == nil
}

func IsNumeric(val interface{}) bool {
	t := reflect.TypeOf(val).Kind()

	if t == reflect.Int || t == reflect.Int8 || t == reflect.Int16 || t == reflect.Int32 || t == reflect.Int64 {
		return true
	} else if t == reflect.Uint || t == reflect.Uint8 || t == reflect.Uint16 || t == reflect.Uint32 || t == reflect.Uint64 {
		return true
	} else if t == reflect.Float32 || t == reflect.Float64 {
		return true
	} else if t == reflect.String {
		if _, err := strconv.ParseFloat(val.(string), 64); err == nil {
			return true
		}
	}

	return false
}

func IsDigit(val interface{}) bool {
	t := reflect.TypeOf(val).Kind()

	if t == reflect.Int || t == reflect.Int8 || t == reflect.Int16 || t == reflect.Int32 || t == reflect.Int64 {
		return true
	} else if t == reflect.Uint || t == reflect.Uint8 || t == reflect.Uint16 || t == reflect.Uint32 || t == reflect.Uint64 {
		return true
	} else if t == reflect.String {
		if _, err := strconv.ParseInt(val.(string), 10, 64); err == nil {
			return true
		}

		if _, err := strconv.ParseUint(val.(string), 10, 64); err == nil {
			return true
		}
	}

	return false
}
