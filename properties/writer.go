package properties

import (
	"bufio"
	"io"
	"strconv"
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
	ascii := strconv.QuoteToASCII(value)
	_, err := w.w.WriteString(key + "=" + ascii[1:len(ascii)-1])
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
