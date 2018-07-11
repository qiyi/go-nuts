package properties

import (
	"strings"
	"testing"
)

func TestUnescape(t *testing.T) {
	s := "k1=\\\\\\n\\t\\=\\#\\!\\:\\ \\u4e2d\\u56fd\n"
	parser, _ := NewParser(strings.NewReader(s))
	node, _ := parser.Next()
	p, ok := node.(*PropertyNode)
	if !ok {
		t.Fail()
	}
	if "\\\n\t=#!: 中国" != p.Value {
		t.Fail()
	}
}
