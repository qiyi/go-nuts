package properties

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

type ContinuationValue struct {
	Values []string
}

// Writer is properties writer
type Writer struct {
	c io.Closer
	w *bufio.Writer
}

func NewWriter(filename string) (*Writer, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &Writer{c: file, w: bufio.NewWriter(file)}, nil
}

func (w *Writer) Flush() error {
	return w.w.Flush()
}

func (w *Writer) Close() error {
	return w.c.Close()
}

func (w *Writer) WriteComment(comment string) error {
	_, err := w.w.WriteString("#" + comment)
	if err != nil {
		return err
	}
	return w.WriteBlank()
}

func (w *Writer) WriteKeyValue(key string, value string) error {
	ascii := strconv.QuoteToASCII(value)
	w.w.WriteString(key + "=" + ascii[1:len(ascii)-1])
	return w.WriteBlank()
}

func (w *Writer) WriteKeyContinuationValue(key string, value ContinuationValue) {
	//TODO
}

func (w *Writer) WriteBlank() error {
	err := w.w.WriteByte('\n')
	if err != nil {
		return err
	}
	return nil
}
