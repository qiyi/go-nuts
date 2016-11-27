package stacks

// StrStack 表示字符串堆栈
type StrStack []string

// Push 向栈内添加元素
func (s StrStack) Push(e string) StrStack {
	return append(s, e)
}

// Empty 返回堆栈是否为空
func (s StrStack) Empty() bool {
	return len(s) == 0
}

// Pop 从栈内取出一个元素
func (s StrStack) Pop() (StrStack, string) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1]
	}
	return s, ""
}

// Peek 查看堆栈的最后一个元素
func (s StrStack) Peek() string {
	l := len(s)
	if l > 0 {
		return s[l-1]
	}
	return ""
}
