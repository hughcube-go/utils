package msslice

import (
	"errors"
	"reflect"
)

// 判断指定元素是否存在数组里面
func InArray(item interface{}, items interface{}) bool {
	return 0 <= Search(item, items)
}

// 在数组里面查找指定元素
func Search(item interface{}, haystack interface{}) int {
	if reflect.TypeOf(haystack).Kind() != reflect.Slice {
		return -1
	}

	items := reflect.ValueOf(haystack)
	for i := 0; i < items.Len(); i++ {
		if reflect.DeepEqual(item, items.Index(i).Interface()) {
			return i
		}
	}

	return -1
}

func In(item interface{}, items ...interface{}) bool {
	return InArray(item, items)
}

func GetElemType(list interface{}, skipPtr bool) (reflect.Type, error) {
	listType := reflect.TypeOf(list)

	if listType.Kind() == reflect.Ptr {
		listType = listType.Elem()
	}

	if listType.Kind() != reflect.Slice {
		return nil, errors.New("List must be of type Slice")
	}

	rowType := listType.Elem()

	if skipPtr && rowType.Kind() == reflect.Ptr {
		rowType = rowType.Elem()
	}

	return rowType, nil
}

func MakeSameTypeValue(list interface{}, len int, cap int) (reflect.Value, error) {
	listType := reflect.TypeOf(list)
	listValue := reflect.ValueOf(list)

	if listType.Kind() == reflect.Ptr {
		listType = listType.Elem()
		listValue = listValue.Elem()
	}

	if listType.Kind() != reflect.Slice {
		return listValue, errors.New("List must be of type Slice")
	}

	slice := reflect.MakeSlice(listType, len, cap)

	return slice, nil
}
