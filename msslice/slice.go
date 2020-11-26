package msslice

import "reflect"

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
