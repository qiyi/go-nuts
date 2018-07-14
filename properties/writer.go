package properties

import (
	"bufio"
	"bytes"
	"io"
)

// ContinuationValue 表示要多行存储的值
type ContinuationValue struct {
	Values []string
}

// Writer 支持 Properties 文件的写入
type Writer struct {
	w *bufio.Writer
}

// NewWriter 从 io.Writer 创建 Properties Writer
func NewWriter(w io.Writer) (*Writer, error) {
	if ww, ok := w.(*bufio.Writer); ok {
		return &Writer{w: ww}, nil
	}
	return &Writer{w: bufio.NewWriter(w)}, nil
}

// Flush 输入 Properties 到 io.Writer
func (w *Writer) Flush() error {
	return w.w.Flush()
}

// WriteComment 写注释
func (w *Writer) WriteComment(comment string) error {
	_, err := w.w.WriteString("#" + comment)
	if err != nil {
		return err
	}
	return w.WriteBlank()
}

// WriteProperty 写一个 Property
func (w *Writer) WriteProperty(key string, value string) error {
	ascii := Escape(value)
	_, err := w.w.WriteString(key + "=" + ascii)
	if err != nil {
		return err
	}
	return w.WriteBlank()
}

// WriteContinuationProperty 写多行值的 Property
func (w *Writer) WriteContinuationProperty(key string, value ContinuationValue) {
	//TODO
}

// WriteBlank 写空白行
func (w *Writer) WriteBlank() error {
	err := w.w.WriteByte('\n')
	if err != nil {
		return err
	}
	return nil
}

const lowerhex = "0123456789abcdef"

// Escape 转义 properties 值
func Escape(s string) string {
	var buf bytes.Buffer
	for _, r := range []rune(s) {
		if r > 61 && r < 127 {
			if r == '\\' {
				buf.WriteRune('\\')
				buf.WriteRune('\\')
				continue
			}
			buf.WriteRune(r)
			continue
		}
		switch r {
		case ' ':
			buf.WriteRune('\\')
			buf.WriteRune(' ')
		case '\t':
			buf.WriteRune('\\')
			buf.WriteRune('t')
		case '\n':
			buf.WriteRune('\\')
			buf.WriteRune('n')
		case '\r':
			buf.WriteRune('\\')
			buf.WriteRune('r')
		case '\f':
			buf.WriteRune('\\')
			buf.WriteRune('f')
		case '=', ':', '#', '!':
			buf.WriteRune('\\')
			buf.WriteRune(r)
		default:
			if r < 0x0020 || r > 0x007e {
				buf.WriteRune('\\')
				buf.WriteRune('u')
				for s := 12; s >= 0; s -= 4 {
					buf.WriteRune(rune(lowerhex[r>>uint(s)&0xF]))
				}
			} else {
				buf.WriteRune(r)
			}
		}
	}
	return buf.String()
}
