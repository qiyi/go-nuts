package properties

import (
	"io"
	"os"
)

// NodeType 表示 Properties 文件里的节点类型
type NodeType int

const (
	CommentType  NodeType = 0 // 注释节点
	BlankType    NodeType = 1 // 空白行
	PropertyType NodeType = 2 // Property 节点
)

// Node 是节点接口
type Node interface {
	Type() NodeType // Type 返回节点类型
}

// CommentNode 是注释节点实现
type CommentNode struct {
	Text string
}

// Type 返回注释节点类型
func (c *CommentNode) Type() NodeType {
	return CommentType
}

// BlankNode 是空行节点实现
type BlankNode struct {
}

// Type 返回空行节点类型
func (b *BlankNode) Type() NodeType {
	return BlankType
}

// PropertyNode 是 Property 节点实现
type PropertyNode struct {
	Key   string
	Value string
}

// Type 返回 Property 节点类型
func (n *PropertyNode) Type() NodeType {
	return PropertyType
}

// Properties 对象
type Properties map[string]string

// Load 从 Reader 加载 Properties
func Load(r io.Reader) (Properties, error) {
	var p Properties = make(map[string]string, 0)
	parser, err := NewParser(r)
	if err != nil {
		return nil, err
	}
	for {
		n, err := parser.Next()
		if err != nil {
			return nil, err
		}
		if n == nil {
			break
		}
		if pn, ok := n.(*PropertyNode); ok {
			p.Put(pn.Key, pn.Value)
		}
	}
	return p, nil
}

// LoadFile 加载 Properties 文件
func LoadFile(name string) (Properties, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Load(file)
}

// Keys 返回 Property 所有的 key, 无序
func (p Properties) Keys() []string {
	keys := make([]string, 0)
	for key, _ := range p {
		keys = append(keys, key)
	}
	return keys
}

// ContainsKey 检查 Properties 是否包含某个 key
func (p Properties) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}

// Get 获取 Properties 的某个 key 的值
func (p Properties) Get(key string) (string, bool) {
	v, ok := p[key]
	return v, ok
}

// GetDefault 获取 Properties 的某个 key 的值, 不存在时返回默认值
func (p Properties) GetDefault(key string, defaultVal string) string {
	if value, ok := p.Get(key); ok {
		return value
	}
	return defaultVal
}

// Put 往 Properties 里放入一个 Key, Value, 如果 key 存在则覆盖
func (p Properties) Put(key string, value string) {
	p[key] = value
}

// Remove 从 Properties 里移除一个 key
func (p Properties) Remove(key string) {
	delete(p, key)
}

// Store 把 Properties 存到某个 Writer 里
func (p Properties) Store(w io.Writer) error {
	writer, err := NewWriter(w)
	if err != nil {
		return err
	}
	for key, value := range p {
		if err := writer.WriteProperty(key, value); err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// StoreFile 把 Properties 存到文件里
func (p Properties) StoreFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return p.Store(file)

}
