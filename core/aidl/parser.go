package aidl

import (
	"fmt"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type DeclarationKind string

const (
	DeclarationInterface  DeclarationKind = "interface"
	DeclarationParcelable DeclarationKind = "parcelable"
	DeclarationEnum       DeclarationKind = "enum"
	DeclarationUnion      DeclarationKind = "union"
)

type MemberKind string

const (
	MemberMethod MemberKind = "method"
	MemberField  MemberKind = "field"
	MemberValue  MemberKind = "value"
)

type SourceRange struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Import struct {
	Name  string      `json:"name"`
	Range SourceRange `json:"range"`
}

type Parameter struct {
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	Direction string      `json:"direction,omitempty"`
	Range     SourceRange `json:"range"`
}

type Member struct {
	Kind       MemberKind  `json:"kind"`
	Name       string      `json:"name"`
	Type       string      `json:"type,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Range      SourceRange `json:"range"`
}

type Declaration struct {
	Kind          DeclarationKind `json:"kind"`
	Name          string          `json:"name"`
	CanonicalName string          `json:"canonicalName"`
	Members       []Member        `json:"members,omitempty"`
	Nested        []Declaration   `json:"nested,omitempty"`
	Range         SourceRange     `json:"range"`
}

type Document struct {
	Path         model.RepoPath     `json:"path"`
	Package      string             `json:"package,omitempty"`
	Imports      []Import           `json:"imports,omitempty"`
	Declarations []Declaration      `json:"declarations,omitempty"`
	Diagnostics  []model.Diagnostic `json:"diagnostics"`
}

func Parse(path model.RepoPath, source string) Document {
	tokens, diagnostics := Lex(source)
	document := Document{Path: path, Imports: make([]Import, 0), Declarations: make([]Declaration, 0), Diagnostics: append(make([]model.Diagnostic, 0), diagnostics...)}
	parser := parser{tokens: tokens, document: &document}
	parser.parse()
	return document
}

type parser struct {
	tokens   []Token
	index    int
	document *Document
}

func (p *parser) parse() {
	for !p.atEOF() {
		p.skipAnnotations()
		if p.peek().Text == "oneway" && p.index+1 < len(p.tokens) && p.tokens[p.index+1].Text == "interface" {
			p.consume()
			p.parseDeclaration()
			continue
		}
		switch p.peek().Text {
		case "package":
			p.parsePackage()
		case "import":
			p.parseImport()
		case "interface", "parcelable", "enum", "union":
			p.parseDeclaration()
		default:
			p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", fmt.Sprintf("unexpected token %q", p.peek().Text), p.peek().Start)
			p.recoverTopLevel()
		}
	}
	if p.document.Package == "" {
		p.addDiagnostic("AIDL_PACKAGE_MISSING", "AIDL document has no package declaration", Position{})
	}
}

func (p *parser) parsePackage() {
	start := p.consume().Start
	name := p.parseQualifiedName()
	if !p.consumeText(";") {
		p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", "package declaration must end with ';'", p.peek().Start)
	}
	if name == "" {
		p.addDiagnostic("AIDL_PACKAGE_MISSING", "package name is missing", start)
		return
	}
	p.document.Package = name
}

func (p *parser) parseImport() {
	start := p.consume().Start
	name := p.parseQualifiedName()
	if name == "" {
		p.addDiagnostic("AIDL_IMPORT_MISSING", "import name is missing", start)
		p.recoverStatement()
		return
	}
	end := p.peek().Start
	if !p.consumeText(";") {
		p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", "import declaration must end with ';'", p.peek().Start)
	} else {
		end = p.previous().End
	}
	p.document.Imports = append(p.document.Imports, Import{Name: name, Range: SourceRange{Start: start, End: end}})
}

func (p *parser) parseDeclaration() {
	declaration := p.parseDeclarationWithPrefix("")
	if declaration != nil {
		p.document.Declarations = append(p.document.Declarations, *declaration)
	}
}

func (p *parser) parseDeclarationWithPrefix(prefix string) *Declaration {
	start := p.consume()
	kind := DeclarationKind(start.Text)
	p.skipAnnotations()
	nameToken := p.peek()
	if nameToken.Kind != TokenIdentifier {
		p.addDiagnostic("AIDL_DECLARATION_MISSING_NAME", "declaration name is missing", nameToken.Start)
		p.recoverStatement()
		return nil
	}
	name := p.consume().Text
	for p.consumeText(".") {
		if p.peek().Kind != TokenIdentifier {
			p.addDiagnostic("AIDL_DECLARATION_MISSING_NAME", "qualified declaration name is incomplete", p.peek().Start)
			break
		}
		name += "." + p.consume().Text
	}
	canonical := prefix
	if canonical == "" {
		canonical = p.document.Package
	}
	if canonical != "" {
		canonical += "."
	}
	canonical += name
	declaration := Declaration{Kind: kind, Name: name, CanonicalName: canonical, Members: make([]Member, 0), Nested: make([]Declaration, 0), Range: SourceRange{Start: start.Start}}
	if p.consumeText(";") {
		declaration.Range.End = p.previous().End
		return &declaration
	}
	for !p.atEOF() && p.peek().Text != "{" && p.peek().Text != ";" {
		p.consume()
	}
	if p.consumeText(";") {
		declaration.Range.End = p.previous().End
		return &declaration
	}
	if !p.consumeText("{") {
		p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", "declaration body must start with '{'", p.peek().Start)
		return nil
	}
	for !p.atEOF() && p.peek().Text != "}" {
		p.skipAnnotations()
		if p.peek().Text == "oneway" && p.index+1 < len(p.tokens) && p.tokens[p.index+1].Text == "interface" {
			p.consume()
		}
		if p.peek().Text == "interface" || p.peek().Text == "parcelable" || p.peek().Text == "enum" || p.peek().Text == "union" {
			if nested := p.parseDeclarationWithPrefix(declaration.CanonicalName); nested != nil {
				declaration.Nested = append(declaration.Nested, *nested)
			}
			continue
		}
		if member, ok := p.parseMember(kind); ok {
			declaration.Members = append(declaration.Members, member)
		}
	}
	if p.consumeText("}") {
		declaration.Range.End = p.previous().End
	} else {
		declaration.Range.End = p.peek().End
	}
	p.consumeText(";")
	return &declaration
}

func (p *parser) parseMember(kind DeclarationKind) (Member, bool) {
	p.skipAnnotations()
	start := p.peek().Start
	if kind == DeclarationEnum {
		name := p.peek()
		if name.Kind != TokenIdentifier {
			p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", "enum value name is missing", start)
			p.recoverStatement()
			return Member{}, false
		}
		p.consume()
		member := Member{Kind: MemberValue, Name: name.Text, Range: SourceRange{Start: start, End: name.End}}
		for !p.atEOF() && p.peek().Text != "," && p.peek().Text != ";" && p.peek().Text != "}" {
			p.consume()
		}
		if p.consumeText(",") || p.consumeText(";") {
			member.Range.End = p.previous().End
		}
		return member, true
	}
	isConstant := p.consumeText("const")
	if isConstant {
		p.skipAnnotations()
	}
	if p.peek().Text == "oneway" || p.peek().Text == "async" {
		p.consume()
	}
	typeName := p.parseType()
	nameToken := p.peek()
	if typeName == "" || nameToken.Kind != TokenIdentifier {
		p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", "member type or name is missing", start)
		p.recoverStatement()
		return Member{}, false
	}
	name := p.consume().Text
	if p.consumeText("(") {
		member := Member{Kind: MemberMethod, Name: name, Type: typeName, Parameters: make([]Parameter, 0), Range: SourceRange{Start: start}}
		for !p.atEOF() && p.peek().Text != ")" {
			p.skipAnnotations()
			parameterStart := p.peek().Start
			direction := ""
			if p.peek().Text == "in" || p.peek().Text == "out" || p.peek().Text == "inout" {
				direction = p.consume().Text
			}
			p.skipAnnotations()
			parameterType := p.parseType()
			parameterName := p.peek()
			if parameterType == "" || parameterName.Kind != TokenIdentifier {
				p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", "method parameter type or name is missing", parameterStart)
				p.recoverParameter()
				continue
			}
			p.consume()
			member.Parameters = append(member.Parameters, Parameter{Name: parameterName.Text, Type: parameterType, Direction: direction, Range: SourceRange{Start: parameterStart, End: parameterName.End}})
			if !p.consumeText(",") && p.peek().Text != ")" {
				p.addDiagnostic("AIDL_UNEXPECTED_TOKEN", "method parameters must be comma separated", p.peek().Start)
				p.recoverParameter()
			}
		}
		if p.consumeText(")") {
			member.Range.End = p.previous().End
		}
		for !p.atEOF() && p.peek().Text != ";" && p.peek().Text != "}" {
			p.consume()
		}
		if p.consumeText(";") {
			member.Range.End = p.previous().End
		}
		return member, true
	}
	member := Member{Kind: MemberField, Name: name, Type: typeName, Range: SourceRange{Start: start, End: nameToken.End}}
	if isConstant {
		member.Kind = MemberValue
	}
	for !p.atEOF() && p.peek().Text != ";" && p.peek().Text != "," && p.peek().Text != "}" {
		p.consume()
	}
	if p.consumeText(";") || p.consumeText(",") {
		member.Range.End = p.previous().End
	}
	return member, true
}

func (p *parser) parseType() string {
	var parts []string
	if p.peek().Kind != TokenIdentifier {
		return ""
	}
	parts = append(parts, p.consume().Text)
	depth := 0
	for !p.atEOF() {
		text := p.peek().Text
		if text == "." {
			parts = append(parts, p.consume().Text)
			if p.peek().Kind == TokenIdentifier {
				parts = append(parts, p.consume().Text)
			}
			continue
		}
		if text == "<" || text == ">" || text == "[" || text == "]" || text == "?" || depth > 0 && text == "," {
			if text == "<" {
				depth++
			}
			if text == ">" && depth > 0 {
				depth--
			}
			parts = append(parts, p.consume().Text)
			continue
		}
		if depth > 0 && p.peek().Kind == TokenIdentifier {
			parts = append(parts, p.consume().Text)
			continue
		}
		break
	}
	return strings.Join(parts, "")
}

func (p *parser) parseQualifiedName() string {
	if p.peek().Kind != TokenIdentifier {
		return ""
	}
	parts := []string{p.consume().Text}
	for p.consumeText(".") {
		if p.peek().Kind != TokenIdentifier {
			break
		}
		parts = append(parts, p.consume().Text)
	}
	return strings.Join(parts, ".")
}

func (p *parser) skipAnnotations() {
	for p.consumeText("@") {
		if p.peek().Kind == TokenIdentifier {
			p.consume()
		}
		if p.consumeText("(") {
			depth := 1
			for !p.atEOF() && depth > 0 {
				switch p.consume().Text {
				case "(":
					depth++
				case ")":
					depth--
				}
			}
		}
	}
}

func (p *parser) recoverStatement() {
	for !p.atEOF() && p.peek().Text != ";" && p.peek().Text != "}" {
		p.consume()
	}
	p.consumeText(";")
}

func (p *parser) recoverParameter() {
	for !p.atEOF() && p.peek().Text != "," && p.peek().Text != ")" {
		p.consume()
	}
	p.consumeText(",")
}

func (p *parser) recoverTopLevel() {
	for !p.atEOF() && p.peek().Text != ";" && p.peek().Text != "package" && p.peek().Text != "import" && p.peek().Text != "interface" && p.peek().Text != "parcelable" && p.peek().Text != "enum" && p.peek().Text != "union" {
		p.consume()
	}
	p.consumeText(";")
}

func (p *parser) addDiagnostic(code, message string, position Position) {
	p.document.Diagnostics = append(p.document.Diagnostics, model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "aidl-parser", Message: message})
}

func (p *parser) peek() Token {
	if p.index >= len(p.tokens) {
		return Token{Kind: TokenEOF}
	}
	return p.tokens[p.index]
}

func (p *parser) previous() Token { return p.tokens[p.index-1] }

func (p *parser) consume() Token {
	token := p.peek()
	if p.index < len(p.tokens) {
		p.index++
	}
	return token
}

func (p *parser) consumeText(text string) bool {
	if p.peek().Text != text {
		return false
	}
	p.consume()
	return true
}

func (p *parser) atEOF() bool { return p.peek().Kind == TokenEOF }
