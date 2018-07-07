package properties

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

// Reader 是对 Properties 文件的解析, 支持保留注释、换行,
// 同时保留所有元素的顺序
type Parser struct {
	data      []rune
	tokenizer *Tokenizer
}

func NewParser(data []rune) *Parser {
	return &Parser{data: data, tokenizer: NewTokenizer(data)}
}

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
		// 非换行、结束、= 令牌之前的字符都收集起来 TODO unicode escape and backslash escape
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
		keyStr, err := strconv.Unquote(`"` + strings.TrimRight(key.String(), "\r\t\f ") + `"`)
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
		valueStr, err := strconv.Unquote(`"` + strings.TrimLeft(value.String(), "\r\t\f ") + `"`)
		if err != nil {
			return nil, err
		}
		return &PropertyNode{Key: keyStr, Value: valueStr}, nil
	}

}
