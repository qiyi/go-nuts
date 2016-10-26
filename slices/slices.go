package slices

import (
	"strconv"
)

// IndexStr 从一个 []string 里查询字符串并返回所在下标
func IndexStr(slice []string, str string) int {
	if len(slice) != 0 {
		for i, s := range slice {
			if s == str {
				return i
			}
		}
	}
	return -1
}

// ContainsStr 测试一个 []string 里是否包含某字符串
func ContainsStr(slice []string, str string) bool {
	return IndexStr(slice, str) != -1
}

// DeleteStrAt 根据下标从 []string 里删除字符串
func DeleteStrAt(slice []string, idx int) ([]string, bool) {
	if idx >= 0 && idx < len(slice) {
		result := append([]string{}, slice[:idx]...)
		result = append(result, slice[idx+1:]...)
		return result, true
	}
	return slice, false
}

// DeleteStr 从 []string 里移除某个字符串
func DeleteStr(slice []string, str string) ([]string, bool) {
	idx := IndexStr(slice, str)
	if idx != -1 {
		return DeleteStrAt(slice, idx)
	}
	return slice, false
}

// AddFirstStr 将一个字符串添加到 []string 的头部
func AddFirstStr(slice []string, str string) []string {
	return append([]string{str}, slice...)
}

// InsertStr 插入一个字符串到指定下标位置
func InsertStr(slice []string, at int, str string) ([]string, bool) {
	if at == len(slice) {
		return append(slice, str), true
	}
	if at >= 0 && at < len(slice) {
		result := append([]string{}, slice[:at]...)
		result = append(result, str)
		result = append(result, slice[at:]...)
		return result, true
	}
	return slice, false
}

// ForeachStr 迭代 []string 执行方法
func ForeachStr(slice []string, f func(int, string)) {
	if slice != nil {
		for i, v := range slice {
			f(i, v)
		}
	}
}

// NewStrSlice 创建一个初始长度为 0 的 []string
func NewStrSlice() []string {
	return make([]string, 0)
}

func NewSlice() []interface{} {
	return make([]interface{}, 0)
}

func AsSlice(slice []string) []interface{} {
	if slice == nil {
		return nil
	}
	result := NewSlice()
	for _, s := range slice {
		result = append(result, s)
	}
	return result

}

func AsStrSlice(slice []interface{}) ([]string, bool) {
	strSlice := make([]string, len(slice))
	allOk := true
	for i, v := range slice {
		var s string
		switch result := v.(type) {
		case string:
			s = result
		case int:
			s = strconv.Itoa(result)
		case bool:
			s = strconv.FormatBool(result)
		default:
			allOk = false
			break
		}
		strSlice[i] = s
	}
	return strSlice, allOk

}
