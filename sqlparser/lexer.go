package sqlparser

import (
	"fmt"
	"unicode"
)

type TokenType uint8

const (
	Select TokenType = iota
	From
	Where
	Group
	By
	As
	Comma
	Undefined
	Eof
	Having
	Number
	Order
	FieldPath
	Identifier
	LeftParenthesis
	RightParenthesis
)

type Token struct {
	Type TokenType
	Value string
}

var (
	keywordMap = map[string]Token {
		"SELECT": Token{
			Type: Select,
			Value: "SELECT",
		},
		"FROM": Token{
			Type: From,
			Value: "FROM",
		},
		"WHERE": Token{
			Type: Where,
			Value: "WHERE",
		},
		"GROUP": Token{
			Type: Group,
			Value: "GROUP",
		},
		"ORDER": Token{
			Type: Order,
			Value: "ORDER",
		},
		"BY": Token{
			Type: By,
			Value: "BY",
		},
		"AS": Token{
			Type: As,
			Value: "AS",
		},
	}

	delimiterMap = map[rune]Token {
		'(': Token{
			Type: LeftParenthesis,
			Value: "(",
		},
		')': Token{
			Type: RightParenthesis,
			Value: ")",
		},
		',': Token{
			Type: Comma,
			Value: ",",
		},
	}

	UndefinedToken = Token {
		Type: Undefined,
		Value: "",
	}

	EofToken = Token {
		Type: Eof,
		Value: "",
	}
)

type Lexer struct {
	currentPosition int
	nextPosition int
	source []rune
}

func NewLexer(source string) *Lexer {

	return &Lexer{
		currentPosition: 0,
		nextPosition: 1,
		source: []rune(source),
	}
}

func (l *Lexer) Tokenize() []Token {

	tokens := []Token{}

	for token := l.NextToken(); token.Type != Eof; token = l.NextToken() {
		if token.Type == Undefined {
			fmt.Errorf("Invalid token %v", token)
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func (l *Lexer) CurrentChar() rune {
	
	if l.currentPosition > len(l.source) - 1 {
		return rune(0)
	}

	return l.source[l.currentPosition]
}

func (l *Lexer) NextChar() rune {

	l.currentPosition += 1
	l.nextPosition += 1

	if l.currentPosition > len(l.source) - 1 {
		return rune(0)
	}

	return l.source[l.currentPosition]
}

func (l *Lexer) PeekChar() rune {

	if l.nextPosition > len(l.source) - 1 {
		return rune(0)
	}

	return l.source[l.nextPosition]
}

func (l *Lexer) SwallowWhitespace() {
	for char := l.CurrentChar(); isWhitespace(char); char = l.CurrentChar() {
		l.NextChar()
	}
}

func (l *Lexer) NextToken() Token {

	l.SwallowWhitespace()
	char := l.CurrentChar()
	
	if isEof(char) {
		return EofToken
	} else if isAlphabetical(char) {
		tokenValue := l.readIdentifier()

		if isKeyword(tokenValue) {
			return keywordMap[tokenValue]
		}

		return Token {
			Type: Identifier,
			Value: tokenValue,
		}
	} else if isNumerical(char) {
		tokenValue := l.readNumber()

		return Token {
			Type: Number,
			Value: tokenValue,
		}

	} else if isDelimiter(char) {
		l.NextChar()
		return delimiterMap[char]
	}
	
	return UndefinedToken
}

func (l *Lexer) readIdentifier() string {
	identifier := string(l.CurrentChar())

	for isAlphabetical(l.NextChar()) {
		identifier = identifier + string(l.CurrentChar())
	}

	return identifier
}

func (l *Lexer) readNumber() string {
	number := string(l.CurrentChar())

	for isNumerical(l.NextChar()) {
		number = number + string(l.CurrentChar())
	}

	return number
}

func isEof(char rune) bool {
	return char == rune(0)
}

func isWhitespace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func isAlphabetical(char rune) bool {
	return unicode.IsLetter(char)
}

func isNumerical(char rune) bool {
	return unicode.IsDigit(char)
}

func isKeyword(key string) bool {

	if _, ok := keywordMap[key]; ok {
		return true
	}

	return false
}

func isDelimiter(char rune) bool {

	if _, ok := delimiterMap[char]; ok {
		return true
	}

	return false
}