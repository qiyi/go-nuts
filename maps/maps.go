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

func GetInt(m map[interface{}]interface{}, key string) (int, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.(int); ok {
            return result, ok
        }
    }
    return 0, false
}

func GetBool(m map[interface{}]interface{}, key string) (bool, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.(bool); ok {
            return result, ok
        }
    }
    return false, false
}

func GetSlice(m map[interface{}]interface{}, key string) ([]interface{}, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.([]interface{}); ok {
            return result, ok
        }
    }
    return nil, false
}

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

func GetMap(m map[interface{}]interface{}, key string) (map[interface{}]interface{}, bool) {
    if value, ok := Get(m, key); ok {
        if result, ok := value.(map[interface{}]interface{}); ok {
            return result, ok
        }
    }
    return nil, false
}

func NewMap() map[interface{}]interface{} {
    return make(map[interface{}]interface{})
}

func NewStrMap() map[string]interface{} {
    return make(map[string]interface{})
}

