package msstruct

import (
	"errors"
	"strconv"
	"strings"
)

type TagSettingItem struct {
	Value    string
	HasValue bool
}

type TagSettings map[string]TagSettingItem

func ParseTag(str string) TagSettings {
	return ParseTagWithSep(str, ";")
}

func ParseTagWithSep(str string, sep string) TagSettings {
	settings := TagSettings{}

	for _, item := range strings.Split(str, sep) {
		if 0 == len(item) {
			continue
		}

		itemSlice := strings.Split(item, ":")
		if 0 == len(itemSlice) {
			continue
		}

		name := strings.ToUpper(itemSlice[0])
		if 0 == len(name) {
			continue
		}

		if 1 == len(itemSlice) {
			settings[name] = TagSettingItem{HasValue: false}
		} else {
			settings[name] = TagSettingItem{Value: strings.Join(itemSlice[1:], ":"), HasValue: true}
		}
	}

	return settings
}

func (t TagSettings) Has(name string) bool {
	name = strings.ToUpper(name)
	_, ok := t[name]
	return ok
}

func (t TagSettings) Get(name string) string {
	name = strings.ToUpper(name)
	if val, ok := t[name]; ok {
		return val.Value
	}
	return ""
}

func (t TagSettings) IsTrue(name string) bool {
	name = strings.ToUpper(name)
	val, ok := t[name]

	if ok && !val.HasValue {
		return true
	}

	if ok {
		value := strings.ToLower(val.Value)
		if "on" == value || "yes" == value || "true" == value || "1" == value {
			return true
		}
	}

	return false
}

func (t TagSettings) GetInt64(name string) (int64, error) {
	name = strings.ToUpper(name)
	val, ok := t[name]

	if ok && !val.HasValue {
		return 0, nil
	}

	if ok {
		if value, err := strconv.ParseInt(val.Value, 10, 64); err == nil {
			return value, nil
		}
	}

	return 0, errors.New("tag value not int")
}

func (t TagSettings) GetInt(name string) (int, error) {
	value, err := t.GetInt64(name)
	return int(value), err
}

func (t TagSettings) GetFloat64(name string) (float64, error) {
	name = strings.ToUpper(name)
	val, ok := t[name]

	if ok && !val.HasValue {
		return 0, nil
	}

	if ok {
		if value, err := strconv.ParseFloat(val.Value, 64); err == nil {
			return value, nil
		}
	}

	return 0, errors.New("tag value not float")
}

func (t TagSettings) GetFloat(name string) (float32, error) {
	value, err := t.GetFloat64(name)
	return float32(value), err
}
