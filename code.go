package stackmachine

import (
	"regexp"
	"strings"
)

const (
	tokenTypeNil = iota
	tokenTypeString
	tokenTypeNumber
	tokenTypeKeyword
	tokenEOF
)

type token struct {
	tokenType int
	value     string
}

type code struct {
	codeText string
	tokens   []token
}

func newCode(codeText string) *code {
	c := &code{
		codeText: codeText,
		tokens:   make([]token, 0),
	}
	c.parseCode()
	return c
}

func (c *code) parseCode() {
	for {
		tokenType, value := c.nextToken()
		if tokenType == tokenEOF {
			break
		}
		c.tokens = append(c.tokens, token{
			tokenType: tokenType,
			value:     value,
		})
	}
}

func (c *code) nextToken() (tokenType int, token string) {
	c.skipWhiteSpaces()
	if len(c.codeText) == 0 {
		return tokenEOF, "EOF"
	}
	switch c.codeText[0] {
	case '"':
		token = c.scanString()
		return tokenTypeString, token
	}
	if isNumber(c.codeText[0]) {
		token = c.scanNumber()
		return tokenTypeNumber, token
	} else if isLetter(c.codeText[0]) {
		token = c.scanIdentifier()
		return tokenTypeKeyword, token
	}
	panic("not valid token")
}

func (c *code) skipWhiteSpaces() {
	for len(c.codeText) > 0 {
		if c.test("\r\n") || c.test("\n\r") {
			c.next(2)
		} else if isNewLine(c.codeText[0]) {
			c.next(1)
		} else if isWhiteSpace(c.codeText[0]) {
			c.next(1)
		} else {
			break
		}
	}
}

func (c *code) next(n int) {
	c.codeText = c.codeText[n:]
}

func (c *code) skipComment() {
	c.next(2)
	if len(c.codeText) > 0 && !isNewLine(c.codeText[0]) {
		c.next(1)
	}
}

func (c *code) test(prefix string) bool {
	return strings.HasPrefix(c.codeText, prefix)
}

var reShortStr = regexp.MustCompile(`(?s)(^"(\\\\|\\"|\\\n|\\z\s*|[^"\n])*")`)

func (c *code) scanString() string {
	if str := reShortStr.FindString(c.codeText); str != "" {
		c.next(len(str))
		str = str[1 : len(str)-1]
		return str
	}
	panic("string not completed")
}

var reNumber = regexp.MustCompile(`(^0[xX][0-9a-fA-F]*)|(^[0-9]+)`)

func (c *code) scanNumber() string {
	if number := reNumber.FindString(c.codeText); number != "" {
		c.next(len(number))
		return number
	}
	panic("not valid number")
}

var reIdentifier = regexp.MustCompile(`^[\d\w]+`)

func (c *code) scanIdentifier() string {
	if identifer := reIdentifier.FindString(c.codeText); identifer != "" {
		c.next(len(identifer))
		return identifer
	}
	panic("identifier not valid")
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
