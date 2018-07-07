package properties

// NodeType 表示 Properties 文件里的节点类型
type NodeType int

const (
	CommentType  NodeType = 0
	BlankType    NodeType = 1
	PropertyType NodeType = 2
)

type Node interface {
	Type() NodeType
}

type CommentNode struct {
	Text string
}

func (c *CommentNode) Type() NodeType {
	return CommentType
}

type BlankNode struct {
}

func (b *BlankNode) Type() NodeType {
	return BlankType
}

type PropertyNode struct {
	Key   string
	Value string
}

func (n *PropertyNode) Type() NodeType {
	return PropertyType
}

// Properties 对象
type Properties map[string]string

func Load(data []byte) (Properties, error) {
	var p Properties = make(map[string]string, 0)
	parser := NewParser([]rune(string(data)))
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

func (p Properties) Keys() []string {
	keys := make([]string, 0)
	for key, _ := range p {
		keys = append(keys, key)
	}
	return keys
}

func (p Properties) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}

func (p Properties) Get(key string) (string, bool) {
	v, ok := p[key]
	return v, ok
}

func (p Properties) GetDefault(key string, defaultVal string) string {
	if value, ok := p.Get(key); ok {
		return value
	}
	return defaultVal
}

func (p Properties) Put(key string, value string) {
	p[key] = value
}

func (p Properties) Remove(key string) {
	delete(p, key)
}

func (p Properties) Store(filename string) error {
	writer, err := NewWriter(filename)
	if err != nil {
		return err
	}
	for key, value := range p {
		writer.WriteKeyValue(key, value)
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return writer.Close()
}
