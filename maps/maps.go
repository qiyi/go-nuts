package maps

// Get 从 map[interface{}]interface{} 里获取值
func Get(m map[interface{}]interface{}, key string) (interface{}, bool) {
    if m != nil {
        if value, ok := m[key]; ok {
            return value, ok
        }
    }
    return nil, false
}

// GetStr 从 map[interface{}]interface{} 里获取 string 值
func GetStr(m map[interface{}]interface{}, key string) (string, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.(string); ok {
            return result, ok
        }
    }
    return "", false
}

// GetInt 从 map[interface{}]interface{} 里获取 int 值
func GetInt(m map[interface{}]interface{}, key string) (int, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.(int); ok {
            return result, ok
        }
    }
    return 0, false
}

// GetBool 从 map[interface{}]interface{} 里获取 bool 值
func GetBool(m map[interface{}]interface{}, key string) (bool, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.(bool); ok {
            return result, ok
        }
    }
    return false, false
}

// GetSlice 从 map[interface{}]interface{} 里获取 []interface{} 值
func GetSlice(m map[interface{}]interface{}, key string) ([]interface{}, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.([]interface{}); ok {
            return result, ok
        }
    }
    return nil, false
}

// GetStrSlice 从 map[interface{}]interface{} 里获取 []string 值
func GetStrSlice(m map[interface{}]interface{}, key string) ([]string, bool) {
    if slice, ok := GetSlice(m, key); ok {
        result := make([]string, 0)
        for _, s := range slice {
            if str, ok := s.(string); ok {
                result = append(result, str)
            }
        }
        return result, true
    }
    return nil, false
}

// GetMapSlice 从 map[interface{}]interface{} 里获取 []map[interface{}]interface{} 值
func GetMapSlice(m map[interface{}]interface{}, key string) ([]map[interface{}]interface{}, bool) {
    if slice, ok := GetSlice(m, key); ok {
        result := make([]map[interface{}]interface{}, 0)
        for _, s := range slice {
            if dic, ok := s.(map[interface{}]interface{}); ok {
                result = append(result, dic)
            }
        }
        return result, true
    }
    return nil, false
}

// GetMap 从 map[interface{}]interface{} 里获取 map[interface{}]interface{} 值
func GetMap(m map[interface{}]interface{}, key string) (map[interface{}]interface{}, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.(map[interface{}]interface{}); ok {
            return result, ok
        }
    }
    return nil, false
}

// NewMap 创建 map[interface{}]interface{}
func NewMap() map[interface{}]interface{} {
    return make(map[interface{}]interface{})
}

// NewStrMap 创建 map[string]interface{}
func NewStrMap() map[string]interface{} {
    return make(map[string]interface{})
}