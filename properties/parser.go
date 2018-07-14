package properties

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Parser 是 Properties 文件的解析器, 支持保留注释、空白行,
// 以及保留所有元素的顺序
type Parser struct {
	data      []rune
	tokenizer *Tokenizer
}

// NewParser 创建新的解析器
func NewParser(r io.Reader) (*Parser, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	data := []rune(string(bytes))
	return &Parser{data: data, tokenizer: NewTokenizer(data)}, nil
}

// Next 获取下一个节点信息
func (p *Parser) Next() (Node, error) {
	token := p.tokenizer.Next()
	if token == nil {
		return nil, nil
	}

	if token.Type == HashToken {
		var comments bytes.Buffer
		for {
			token = p.tokenizer.Next()
			if token == nil || token.Type == NewLineToken {
				return &CommentNode{Text: comments.String()}, nil
			} else if token.Type == ContinueLineToken {
				comments.WriteString(string(p.data[token.Offset : token.Offset+token.Length]))
				return &CommentNode{Text: comments.String()}, nil
			} else {
				comments.WriteString(string(p.data[token.Offset : token.Offset+token.Length]))
			}
		}
	} else if token.Type == NewLineToken {
		return &BlankNode{}, nil
	} else if token.Type == ContinueLineToken {
		return nil, errors.New("invalid continuation characters")
	} else if token.Type == EqualsToken {
		return nil, errors.New("invalid = at the beginning of a line")
	} else {
		// 非换行、结束、= 令牌之前的字符都收集起来
		// get key
		var key bytes.Buffer
		key.WriteString(string(p.data[token.Offset : token.Offset+token.Length]))
		equalsExist := false
		for {
			token = p.tokenizer.Next()
			if token == nil || token.Type == NewLineToken {
				break
			} else if token.Type == EqualsToken {
				equalsExist = true
				break
			} else if token.Type == HashToken {
				return nil, errors.New("invalid hash character in key")
			} else if token.Type == ContinueLineToken {
				continue
			}
			key.WriteString(string(p.data[token.Offset : token.Offset+token.Length]))
		}
		if !equalsExist {
			return nil, errors.New("no value provide for " + key.String())
		}
		keyStr, err := Unescape(strings.TrimRight(key.String(), "\r\t\f "))
		if err != nil {
			return nil, err
		}

		// get value
		var value bytes.Buffer
		for {
			token = p.tokenizer.Next()
			if token == nil || token.Type == NewLineToken {
				break
			}
			if token.Type == ContinueLineToken {
				continue
			}
			value.WriteString(string(p.data[token.Offset : token.Offset+token.Length]))
		}
		valueStr, err := Unescape(strings.TrimLeft(value.String(), "\r\t\f "))
		if err != nil {
			return nil, err
		}
		return &PropertyNode{Key: keyStr, Value: valueStr}, nil
	}
}

// Unescape 反转义 Properties 值
func Unescape(s string) (string, error) {
	r := []rune(s)
	var buf bytes.Buffer
	escaped := false
	position := 0
	for position < len(r) {
		c := r[position]
		position++
		if c == '\\' && !escaped {
			escaped = true
			continue
		}
		if !escaped {
			buf.WriteRune(c)
		} else {
			switch c {
			case 'r':
				buf.WriteRune('\r')
			case 'n':
				buf.WriteRune('\n')
			case 'f':
				buf.WriteRune('\f')
			case 't':
				buf.WriteRune('\t')
			case 'u':
				if position+4 > len(s) {
					return "", strconv.ErrSyntax
				}
				var v rune
				for i := 0; i < 4; i++ {
					x, ok := unhex(byte(r[position]))
					if !ok {
						return "", strconv.ErrSyntax
					}
					v = v<<4 | x
					position++

				}
				buf.WriteRune(v)
			default:
				buf.WriteRune(c)

			}
		}
		escaped = false
	}
	return buf.String(), nil
}

func unhex(b byte) (v rune, ok bool) {
	c := rune(b)
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}
	return
}
