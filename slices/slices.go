package slices

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
        return append(slice[:idx], slice[idx + 1:]...), true
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



