package mscheck

import (
	"net/url"
	"regexp"
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
