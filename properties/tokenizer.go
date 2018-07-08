package properties

// TODO 不支持 ! 开头的注释

// TokenType 表示令牌类型
type TokenType int

const (
	HashToken         TokenType = 0 // HashToken 表示 #
	TextToken         TokenType = 1 // TextToken 表示文本
	EqualsToken       TokenType = 2 // EqualsToken 表示  =
	NewLineToken      TokenType = 3 // NewLineToken 表示换行
	ContinueLineToken TokenType = 4 // ContinueLineToken 表示行连接符号
)

// Token 表示令牌信息
type Token struct {
	Type   TokenType
	Offset int
	Length int
}

// Tokenizer 是令牌解析器
type Tokenizer struct {
	data          []rune
	position      int
	tokenPosition int
	last          *Token
}

// NewTokenizer 新建令牌解析器
func NewTokenizer(data []rune) *Tokenizer {
	return &Tokenizer{data: data}
}

func (t *Tokenizer) newToken(tokenType TokenType) *Token {
	return &Token{Type: tokenType, Offset: t.tokenPosition, Length: t.position - t.tokenPosition}
}

func (t *Tokenizer) skipWhiteSpaces() {
	for {
		c := t.data[t.position]
		if c == '\r' || c == '\t' || c == ' ' || c == '\f' {
			t.position++
			continue
		}
		return
	}
}

// Next 获取下一个令牌，结束了返回 nil
func (t *Tokenizer) Next() *Token {
	if t.position >= len(t.data) {
		return nil
	}
	if t.last == nil || t.last.Type == NewLineToken || t.last.Type == ContinueLineToken {
		t.skipWhiteSpaces()
	}
	t.tokenPosition = t.position
	c := t.data[t.position]
	if c == '#' {
		t.position++
		t.last = t.newToken(HashToken)
		return t.last
	} else if c == '=' {
		t.position++
		t.last = t.newToken(EqualsToken)
	} else if lineSepLen := t.isNewLine(t.position); lineSepLen > 0 {
		t.position = t.position + lineSepLen
		t.last = t.newToken(NewLineToken)
	} else {
		if c == '\\' {
			lineSepLen := t.isNewLine(t.position + 1)
			if lineSepLen > 0 {
				t.position++
				t.last = t.newToken(ContinueLineToken)
				t.position = t.position + lineSepLen
				return t.last
			}
		}
		t.last = t.nextStr()
	}
	return t.last
}

func (t *Tokenizer) isNewLine(offset int) int {
	if offset < len(t.data) && t.data[offset] == '\n' {
		return 1
	} else if offset+1 < len(t.data) && t.data[offset] == '\r' && t.data[offset+1] != '\n' {
		return 1
	} else if offset+1 < len(t.data) && t.data[offset] == '\r' && t.data[offset+1] == '\n' {
		return 2
	}
	return -1
}

func (t *Tokenizer) nextStr() *Token {
	escaped := false
	for t.position < len(t.data) {
		c := t.data[t.position]
		if c == '\\' && !escaped && t.isNewLine(t.position+1) > 0 {
			return t.newToken(TextToken)
		} else if c == '=' && !escaped {
			return t.newToken(TextToken)
		} else if c == '\n' && !escaped {
			return t.newToken(TextToken)
		} else if c == '\r' && t.position+1 <= len(t.data) && t.data[t.position+1] == '\n' && !escaped {
			return t.newToken(TextToken)
		}
		if c == '\\' {
			escaped = !escaped
		} else {
			escaped = false
		}
		t.position++
	}
	return t.newToken(TextToken)
}
