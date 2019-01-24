package stackmachine

import (
	"regexp"
	"strings"
)

const (
	tokenTypeNil = iota
	tokenTypeString
	tokenTypeNumber
	tokenTypeBoolean
	tokenTypeKeyword
	tokenTypeEOF
	tokenTypeLabel
)

type token struct {
	tokenType int
	value     string
}

type lexer struct {
	codeText string
	tokens   []token
}

func newLexer(codeText string) *lexer {
	l := &lexer{
		codeText: codeText,
		tokens:   make([]token, 0),
	}
	l.init()
	return l
}

func (l *lexer) init() {
	for {
		tokenType, value := l.nextToken()
		if tokenType == tokenTypeEOF {
			break
		}
		l.tokens = append(l.tokens, token{
			tokenType: tokenType,
			value:     value,
		})
	}
}

func (l *lexer) nextToken() (tokenType int, token string) {
	l.skipWhiteSpaces()
	if len(l.codeText) == 0 {
		return tokenTypeEOF, "EOF"
	}
	switch l.codeText[0] {
	case '"':
		token = l.scanString()
		return tokenTypeString, token
	case '#':
		token = l.scanLabel()
		return tokenTypeLabel, token
	}
	if isNumber(l.codeText[0]) {
		token = l.scanNumber()
		return tokenTypeNumber, token
	} else if isLetter(l.codeText[0]) {
		token = l.scanIdentifier()
		if strings.ToLower(token) == "true" || strings.ToLower(token) == "false" {
			return tokenTypeBoolean, strings.ToLower(token)
		}
		return tokenTypeKeyword, token
	}
	panic("not valid token")
}

func (l *lexer) skipWhiteSpaces() {
	for len(l.codeText) > 0 {
		if l.test("\r\n") || l.test("\n\r") {
			l.next(2)
		} else if isNewLine(l.codeText[0]) {
			l.next(1)
		} else if isWhiteSpace(l.codeText[0]) {
			l.next(1)
		} else {
			break
		}
	}
}

func (l *lexer) next(n int) {
	l.codeText = l.codeText[n:]
}

func (l *lexer) skipComment() {
	l.next(2)
	if len(l.codeText) > 0 && !isNewLine(l.codeText[0]) {
		l.next(1)
	}
}

func (l *lexer) test(prefix string) bool {
	return strings.HasPrefix(l.codeText, prefix)
}

var reShortStr = regexp.MustCompile(`(?s)(^"(\\\\|\\"|\\\n|\\z\s*|[^"\n])*")`)

func (l *lexer) scanString() string {
	if str := reShortStr.FindString(l.codeText); str != "" {
		l.next(len(str))
		str = str[1 : len(str)-1]
		return str
	}
	panic("string not completed")
}

var reNumber = regexp.MustCompile(`(^0[xX][0-9a-fA-F]*)|(^[0-9]+)`)

func (l *lexer) scanNumber() string {
	if number := reNumber.FindString(l.codeText); number != "" {
		l.next(len(number))
		return number
	}
	panic("not valid number")
}

var reIdentifier = regexp.MustCompile(`^[\d\w]+`)

func (l *lexer) scanIdentifier() string {
	if identifer := reIdentifier.FindString(l.codeText); identifer != "" {
		l.next(len(identifer))
		return identifer
	}
	panic("identifier not valid")
}

var reLabel = regexp.MustCompile(`^#[\d\w]+`)

func (l *lexer) scanLabel() string {
	if label := reLabel.FindString(l.codeText); label != "" {
		l.next(len(label))
		return label
	}
	panic("label not valid")
}

func isWhiteSpace(c byte) bool {
	switch c {
	case '\t', '\v', '\f', '\r', ' ':
		return true
	}
	return false
}

func isNewLine(c byte) bool {
	return c == '\r' || c == '\n'
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func parseCode(l *lexer) *code {
	c := newCode()
	pc := 0
	for {
		if pc >= len(l.tokens) {
			return c
		}
		t := l.tokens[pc]
		switch t.tokenType {
		case tokenTypeKeyword:
			if ins, ok := keywords[t.value]; ok {
				if ins.arity > 0 {
					tokens := make([]token, 0)
					arity := 0
					for arity < ins.arity {
						pc++
						arity++
						tk := l.tokens[pc]
						tokens = append(tokens, tk)
					}
					ins.values = tokens
				}
				c.instructionList = append(c.instructionList, ins)
				if pc+1 < len(l.tokens) && l.tokens[pc+1].tokenType == tokenTypeLabel {
					pc++
					c.labelMap[l.tokens[pc].value] = len(c.instructionList) - 1
				}
				pc++
			}
		}
	}
}
