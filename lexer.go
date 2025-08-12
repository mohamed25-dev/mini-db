package main

import (
	"strings"
)

type TokenType string

const (
	EOF TokenType = ""

	L_PAREN     TokenType = "("
	R_PAREN     TokenType = ")"
	COMMA       TokenType = ","
	MINUS       TokenType = "-"
	PLUS        TokenType = "+"
	ASTERISK    TokenType = "*"
	SEMICOLON   TokenType = ";"
	EQUAL       TokenType = "="
	GT          TokenType = ">"
	LT          TokenType = "<"
	GT_OR_EQUAL TokenType = ">="
	LT_OR_EQUAL TokenType = "<="

	// keywords
	SELECT TokenType = "SELECT"
	FROM   TokenType = "FROM"
	WHERE  TokenType = "WHERE"
	CREATE TokenType = "CREATE"
	TABLE  TokenType = "table"
	INSERT TokenType = "INSERT"
	INTO   TokenType = "INTO"
	VALUES TokenType = "VALUES"

	//  Identifiers + literals
	IDENT  TokenType = "IDENT"
	STRING TokenType = "STRING"
	NUMBER TokenType = "NUMBER"
)

var keywordsMap = map[string]TokenType{
	"SELECT": SELECT,
	"FROM":   FROM,
	"WHERE":  WHERE,
	"CREATE": CREATE,
	"TABLE":  TABLE,
	"INSERT": INSERT,
	"INTO":   INTO,
	"VALUES": VALUES,
}

type Token struct {
	Type    TokenType
	Literal string
}

func NewScanner(input string) *Scanner {
	s := &Scanner{
		input: input,
	}

	s.readChar()
	return s
}

type Scanner struct {
	input        string
	ch           byte // current character under examination
	psosition    int
	readPosition int
}

func (s *Scanner) getNextToken() Token {
	var token Token

	s.skipWhiteSpaces()

	switch s.ch {
	case '(':
		token = Token{
			Type:    L_PAREN,
			Literal: "(",
		}

	case ')':
		token = Token{
			Type:    R_PAREN,
			Literal: ")",
		}

	case '-':
		token = Token{
			Type:    MINUS,
			Literal: "-",
		}

	case '+':
		token = Token{
			Type:    PLUS,
			Literal: "+",
		}

	case '*':
		token = Token{
			Type:    ASTERISK,
			Literal: "*",
		}

	case ';':
		token = Token{
			Type:    SEMICOLON,
			Literal: ";",
		}

	case '=':
		token = Token{
			Type:    EQUAL,
			Literal: "=",
		}

	case '>':
		if s.peekChar() == '=' {
			s.readChar()

			token = Token{
				Type:    GT_OR_EQUAL,
				Literal: ">=",
			}
		} else {
			token = Token{
				Type:    GT,
				Literal: ">",
			}
		}

	case '<':
		if s.peekChar() == '=' {
			s.readChar()

			token = Token{
				Type:    LT_OR_EQUAL,
				Literal: "<=",
			}
		} else {
			token = Token{
				Type:    LT,
				Literal: "<",
			}
		}

	case '\'', '"':
		literal := s.readString(s.ch)
		token = Token{
			Type:    STRING,
			Literal: literal,
		}

	case 0:
		token = Token{
			Type:    EOF,
			Literal: "",
		}

	default:
		if isLetter(s.ch) {
			ident := s.readIdentifier()
			tokenType, ok := keywordsMap[strings.ToUpper(ident)]
			if !ok {
				token = Token{
					Type:    IDENT,
					Literal: ident,
				}
			} else {
				token = Token{
					Type:    tokenType,
					Literal: ident,
				}
			}

		} else if isDigit(s.ch) {
			ident := s.readNumber()
			token = Token{
				Type:    NUMBER,
				Literal: ident,
			}
		} else {
			token = Token{
				Type:    EOF,
				Literal: "",
			}
		}
	}

	s.readChar()
	return token
}

func (s *Scanner) readIdentifier() string {
	position := s.psosition
	for isDigit(s.peekChar()) || isLetter(s.peekChar()) {
		s.readChar()
	}

	return s.input[position : s.psosition+1]
}

func (s *Scanner) readString(quote byte) string {
	position := s.psosition + 1 // skip the quote
	for {
		s.readChar()
		if s.ch == quote || s.ch == 0 {
			break
		}
	}

	return s.input[position:s.psosition]
}

func (s *Scanner) readNumber() string {
	position := s.psosition
	for isDigit(s.peekChar()) {
		s.readChar()
	}

	return s.input[position : s.psosition+1]
}

func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input) {
		s.ch = 0
	} else {
		s.ch = s.input[s.readPosition]
	}

	s.psosition = s.readPosition
	s.readPosition++
}

func (s *Scanner) peekChar() byte {
	if s.readPosition >= len(s.input) {
		return 0
	} else {
		return s.input[s.readPosition]
	}
}

func (s *Scanner) skipWhiteSpaces() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.readChar()
	}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
