package properties

import (
	"strings"
	"testing"
)

// TestEscape 测试转义
func TestEscape(t *testing.T) {
	var sb strings.Builder
	writer, _ := NewWriter(&sb)
	writer.WriteProperty("k1", "\\\n\t=#!: 中国")
	writer.Flush()
	if "k1=\\\\\\n\\t\\=\\#\\!\\:\\ \\u4e2d\\u56fd\n" != sb.String() {
		t.Fail()
	}
}
