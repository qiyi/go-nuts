package maps

import (
	"errors"
	"strings"

	"github.com/qiyi/go-nuts/slices"
)

// Get 从 map[interface{}]interface{} 里获取值
func Get(m map[interface{}]interface{}, key interface{}) (interface{}, bool) {
	if m != nil {
		if value, ok := m[key]; ok {
			return value, ok
		}
	}
	return nil, false
}

// GetStr 从 map[interface{}]interface{} 里获取 string 值
func GetStr(m map[interface{}]interface{}, key interface{}) (string, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.(string); ok {
			return result, ok
		}
	}
	return "", false
}

// GetInt 从 map[interface{}]interface{} 里获取 int 值
func GetInt(m map[interface{}]interface{}, key interface{}) (int, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.(int); ok {
			return result, ok
		}
	}
	return 0, false
}

// GetBool 从 map[interface{}]interface{} 里获取 bool 值
func GetBool(m map[interface{}]interface{}, key interface{}) (bool, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.(bool); ok {
			return result, ok
		}
	}
	return false, false
}

// GetSlice 从 map[interface{}]interface{} 里获取 []interface{} 值
func GetSlice(m map[interface{}]interface{}, key interface{}) ([]interface{}, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.([]interface{}); ok {
			return result, true
		}
	}
	return nil, false
}

// GetStrSlice 从 map[interface{}]interface{} 里获取 []string 值
func GetStrSlice(m map[interface{}]interface{}, key interface{}) ([]string, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.([]string); ok {
			return result, true
		}
	}
	return nil, false
}

// GetMap 从 map[interface{}]interface{} 里获取 map[interface{}]interface{} 值
func GetMap(m map[interface{}]interface{}, key interface{}) (map[interface{}]interface{}, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.(map[interface{}]interface{}); ok {
			return result, true
		}
	}
	return nil, false
}

// GetStrMap 从 map[interface{}]interface{} 里获取 map[string]interface{} 值
func GetStrMap(m map[interface{}]interface{}, key interface{}) (map[string]interface{}, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.(map[string]interface{}); ok {
			return result, ok
		}
	}
	return nil, false
}

// GetMapSlice 从 map[interface{}]interface{} 里获取 []map[interface{}]interface{} 值
func GetMapSlice(m map[interface{}]interface{}, key interface{}) ([]map[interface{}]interface{}, bool) {
	if value, ok := Get(m, key); ok {
		if result, ok := value.([]map[interface{}]interface{}); ok {
			return result, true
		}
	}
	return nil, false
}

// CloneMap 克隆 map[interface{}]interface{} 到一个新的 map[interface{}]interface
func CloneMap(m map[interface{}]interface{}) map[interface{}]interface{} {
	result := NewMap()
	for key, value := range m {
		result[key] = value
	}
	return result
}

// CloneStrMap 克隆 map[string]interface{} 到一个新的 map[string]interface
func CloneStrMap(m map[string]interface{}) map[string]interface{} {
	result := NewStrMap()
	for key, value := range m {
		result[key] = value
	}
	return result
}

// AsMap 将 map[string]interface{} 转换成 map[interface{}]interface{}
func AsMap(m map[string]interface{}) map[interface{}]interface{} {
	result := NewMap()
	for key, value := range m {
		result[key] = value
	}
	return result
}

// AsStrMap 将 map[interface{}]interface{} 转换成 map[string]interface{}
func AsStrMap(m map[interface{}]interface{}) (map[string]interface{}, bool) {
	result := NewStrMap()
	allOk := true
	for key, value := range m {
		if k, ok := key.(string); ok {
			result[k] = value
		} else {
			allOk = false
		}
	}
	return result, allOk
}

// NewMap 创建 map[interface{}]interface{}
func NewMap() map[interface{}]interface{} {
	return make(map[interface{}]interface{})
}

// NewStrMap 创建 map[string]interface{}
func NewStrMap() map[string]interface{} {
	return make(map[string]interface{})
}

// visitValue 根据 path 访问 map[interface{}]interface 获取 interface{} 值
func visitValue(m map[interface{}]interface{}, paths []interface{}) (interface{}, error) {
	if m == nil {
		return nil, errors.New("Nil value")
	}

	if paths == nil {
		return m, nil
	}

	var visitErr = func(p []interface{}) error {
		if ps, ok := slices.AsStrSlice(p); ok {
			return errors.New("Visit failed: " + strings.Join(ps, ","))
		}
		return errors.New("Visit failed")
	}

	var node interface{} = m
	visited := make([]interface{}, 0)
	for _, path := range paths {
		visited = append(visited, path)
		switch n := node.(type) {
		case map[interface{}]interface{}:
			if result, ok := Get(n, path); ok {
				node = result
			} else {
				return nil, visitErr(visited)
			}
		case map[string]interface{}:
			if result, ok := Get(AsMap(n), path); ok {
				node = result
			} else {
				return nil, visitErr(visited)
			}
		case []map[interface{}]interface{}:
			if idx, ok := path.(int); ok {
				if len(n) > idx {
					node = n[idx]
					continue
				}
			}
			return nil, visitErr(visited)
		case []map[string]interface{}:
			if idx, ok := path.(int); ok {
				if len(n) > idx {
					node = n[idx]
					continue
				}
			}
			return nil, visitErr(visited)
		case []interface{}:
			if idx, ok := path.(int); ok {
				if len(n) > idx {
					node = n[idx]
					continue
				}
			}
			return nil, visitErr(visited)
		case []string:
			if idx, ok := path.(int); ok {
				if len(n) > idx {
					node = n[idx]
					continue
				}
			}
			return nil, visitErr(visited)
		default:
			return nil, visitErr(visited)
		}
	}
	return node, nil
}

// Visit 根据 path 访问 map[interface{}]interface 获取 interface{} 值
func Visit(m map[interface{}]interface{}, paths ...interface{}) (interface{}, error) {
	return visitValue(m, paths)
}

// VisitStr 根据 path 访问 map[interface{}]interface 获取 string 值
func VisitStr(m map[interface{}]interface{}, paths ...interface{}) (string, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.(string); ok {
			return result, nil
		} else {
			return "", errors.New("Type assert error")
		}
	} else {
		return "", err
	}
}

// VisitInt 根据 path 访问 map[interface{}]interface 获取 int 值
func VisitInt(m map[interface{}]interface{}, paths ...interface{}) (int, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.(int); ok {
			return result, nil
		} else {
			return 0, errors.New("Type assert error")
		}
	} else {
		return 0, err
	}
}

// VisitBool 根据 path 访问 map[interface{}]interface 获取 bool 值
func VisitBool(m map[interface{}]interface{}, paths ...interface{}) (bool, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.(bool); ok {
			return result, nil
		} else {
			return false, errors.New("Type assert error")
		}
	} else {
		return false, err
	}
}

func visitSliceValue(m map[interface{}]interface{}, paths []interface{}) ([]interface{}, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.([]interface{}); ok {
			return result, nil
		}
		if result, ok := value.([]string); ok {
			return slices.AsSlice(result), nil
		}
		if ms, ok := value.([]map[string]interface{}); ok {
			result := make([]interface{}, len(ms))
			for i, s := range ms {
				result[i] = s
			}
			return result, nil
		}
		if ms, ok := value.([]map[interface{}]interface{}); ok {
			result := make([]interface{}, len(ms))
			for i, s := range ms {
				result[i] = s
			}
			return result, nil
		}
		return nil, errors.New("Type assert error")
	} else {
		return nil, err
	}
}

// VisitSlice 根据 path 访问 map[interface{}]interface 获取 []interface{} 值
func VisitSlice(m map[interface{}]interface{}, paths ...interface{}) ([]interface{}, error) {
	return visitSliceValue(m, paths)
}

// VisitStrSlice 根据 path 访问 map[interface{}]interface 获取 []string 值
func VisitStrSlice(m map[interface{}]interface{}, paths ...interface{}) ([]string, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.([]string); ok {
			return result, nil
		}
		return nil, errors.New("Type assert error")
	} else {
		return nil, err
	}
}

// VisitMap 根据 path 访问 map[interface{}]interface 获取 map[interface{}]interface{} 值
func VisitMap(m map[interface{}]interface{}, paths ...interface{}) (map[interface{}]interface{}, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.(map[interface{}]interface{}); ok {
			return result, nil
		}
		if result, ok := value.(map[string]interface{}); ok {
			return AsMap(result), nil
		}
		return nil, errors.New("Type assert error")
	} else {
		return nil, err
	}
}

// VisitStrMap 根据 path 访问 map[interface{}]interface 获取 map[string]interface{} 值
func VisitStrMap(m map[interface{}]interface{}, paths ...interface{}) (map[string]interface{}, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.(map[string]interface{}); ok {
			return result, nil
		}
		return nil, errors.New("Type assert error")
	} else {
		return nil, err
	}
}

// VisitMapSlice 根据 path 访问 map[interface{}]interface{} 获取 []map[interface{}]interface{} 值
func VisitMapSlice(m map[interface{}]interface{}, paths ...interface{}) ([]map[interface{}]interface{}, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.([]map[interface{}]interface{}); ok {
			return result, nil
		}
		if ms, ok := value.([]map[string]interface{}); ok {
			result := make([]map[interface{}]interface{}, len(ms))
			for i, v := range ms {
				result[i] = AsMap(v)
			}
			return result, nil
		}
	}

	if slice, err := visitSliceValue(m, paths); err == nil {
		result := make([]map[interface{}]interface{}, 0)
		for _, s := range slice {
			if dic, ok := s.(map[interface{}]interface{}); ok {
				result = append(result, dic)
			}
			if dic, ok := s.(map[string]interface{}); ok {
				result = append(result, AsMap(dic))
			}
		}
		return result, nil
	} else {
		return nil, err
	}
}

// VisitStrMapSlice 根据 path 访问 map[interface{}]interface{} 获取 []map[string]interface{} 值
func VisitStrMapSlice(m map[interface{}]interface{}, paths ...interface{}) ([]map[string]interface{}, error) {
	if value, err := visitValue(m, paths); err == nil {
		if result, ok := value.([]map[string]interface{}); ok {
			return result, nil
		}
	}
	slice, err := visitSliceValue(m, paths)
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0)
	for _, s := range slice {
		if dic, ok := s.(map[string]interface{}); ok {
			result = append(result, dic)
		}
	}
	return result, nil
}
