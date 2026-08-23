package aidl

import "github.com/i-rtfsc/as-aosp/core/model"

type TokenKind string

const (
	TokenIdentifier TokenKind = "identifier"
	TokenString     TokenKind = "string"
	TokenNumber     TokenKind = "number"
	TokenSymbol     TokenKind = "symbol"
	TokenEOF        TokenKind = "eof"
)

type Position struct {
	Offset int `json:"offset"`
	Line   int `json:"line"`
	Column int `json:"column"`
}

type Token struct {
	Kind  TokenKind `json:"kind"`
	Text  string    `json:"text"`
	Start Position  `json:"start"`
	End   Position  `json:"end"`
}

func Lex(source string) ([]Token, []model.Diagnostic) {
	tokens := make([]Token, 0)
	diagnostics := make([]model.Diagnostic, 0)
	position := Position{Line: 1, Column: 1}
	for position.Offset < len(source) {
		ch := source[position.Offset]
		if isSpace(ch) {
			advance(&position, ch)
			continue
		}
		if ch == '/' && position.Offset+1 < len(source) && source[position.Offset+1] == '/' {
			advance(&position, source[position.Offset])
			advance(&position, source[position.Offset])
			for position.Offset < len(source) && source[position.Offset] != '\n' {
				advance(&position, source[position.Offset])
			}
			continue
		}
		if ch == '/' && position.Offset+1 < len(source) && source[position.Offset+1] == '*' {
			start := position
			advance(&position, source[position.Offset])
			advance(&position, source[position.Offset])
			closed := false
			for position.Offset < len(source) {
				if source[position.Offset] == '*' && position.Offset+1 < len(source) && source[position.Offset+1] == '/' {
					advance(&position, source[position.Offset])
					advance(&position, source[position.Offset])
					closed = true
					break
				}
				advance(&position, source[position.Offset])
			}
			if !closed {
				diagnostics = append(diagnostics, lexerDiagnostic("AIDL_UNTERMINATED_COMMENT", "unterminated block comment", start))
			}
			continue
		}
		start := position
		if isIdentifierStart(ch) {
			for position.Offset < len(source) && isIdentifierPart(source[position.Offset]) {
				advance(&position, source[position.Offset])
			}
			tokens = append(tokens, Token{Kind: TokenIdentifier, Text: source[start.Offset:position.Offset], Start: start, End: position})
			continue
		}
		if ch >= '0' && ch <= '9' {
			for position.Offset < len(source) && isNumberPart(source[position.Offset]) {
				advance(&position, source[position.Offset])
			}
			tokens = append(tokens, Token{Kind: TokenNumber, Text: source[start.Offset:position.Offset], Start: start, End: position})
			continue
		}
		if ch == '"' {
			advance(&position, ch)
			closed := false
			for position.Offset < len(source) {
				ch = source[position.Offset]
				advance(&position, ch)
				if ch == '\\' && position.Offset < len(source) {
					advance(&position, source[position.Offset])
					continue
				}
				if ch == '"' {
					closed = true
					break
				}
			}
			if !closed {
				diagnostics = append(diagnostics, lexerDiagnostic("AIDL_UNTERMINATED_STRING", "unterminated string", start))
			}
			tokens = append(tokens, Token{Kind: TokenString, Text: source[start.Offset:position.Offset], Start: start, End: position})
			continue
		}
		if isSymbol(ch) {
			advance(&position, ch)
			tokens = append(tokens, Token{Kind: TokenSymbol, Text: string(ch), Start: start, End: position})
			continue
		}
		diagnostics = append(diagnostics, lexerDiagnostic("AIDL_UNKNOWN_CHARACTER", "unknown character", start))
		advance(&position, ch)
	}
	tokens = append(tokens, Token{Kind: TokenEOF, Start: position, End: position})
	return tokens, diagnostics
}

func advance(position *Position, ch byte) {
	position.Offset++
	if ch == '\n' {
		position.Line++
		position.Column = 1
	} else {
		position.Column++
	}
}

func isSpace(ch byte) bool { return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' }

func isIdentifierStart(ch byte) bool {
	return ch == '_' || ch == '$' || ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z'
}

func isIdentifierPart(ch byte) bool { return isIdentifierStart(ch) || ch >= '0' && ch <= '9' }

func isNumberPart(ch byte) bool {
	return ch >= '0' && ch <= '9' || ch == 'x' || ch == 'X' || ch >= 'a' && ch <= 'f' || ch >= 'A' && ch <= 'F'
}

func isSymbol(ch byte) bool {
	switch ch {
	case '@', '.', '{', '}', '(', ')', '[', ']', '<', '>', ',', ';', '=', ':', '?', '+', '-', '*', '&', '|':
		return true
	default:
		return false
	}
}

func lexerDiagnostic(code, message string, position Position) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "aidl-lexer", Message: message}
}
