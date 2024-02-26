package main

import (
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

type CharSlice []rune

type ParserFunction func(CharSlice, *int) *Token

type Token struct {
	Text         string    `json:"text"`
	Type         TokenType `json:"type"`
	Start        int       `json:"start"`
	End          int       `json:"end"`
	LineNumber   int       `json:"lineNumber"`
	ColumnNumber int       `json:"columnNumber"`
}

type TokenIterator func() *Token

type TokenType int

const (
	AlphaNumeric = iota
	BackticksQuoted
	DoubleQuoted
	Numeric
	Operator
	SingleQuoted
	Symbol
)

type ValidationFunction func(CharSlice, *int) bool

func Parse(text string) TokenIterator {
	pos, inputs := 0, CharSlice(text)
	return func() *Token {
		// skip-over whitespace
		for isWhitespace(inputs, &pos) {
			pos++
		}

		// get the next token
		return nextToken(inputs, &pos)
	}
}

func ParseFully(text string) ([]Token, error) {
	var tokens []Token
	inputs := CharSlice(text)
	for pos := 0; pos < len(inputs); {
		// skip-over whitespace
		for isWhitespace(inputs, &pos) {
			pos++
		}

		// append the next token
		if hasMore(inputs, &pos) {
			token := nextToken(inputs, &pos)
			tokens = append(tokens, *token)
		}
	}
	return tokens, nil
}

func (token Token) GetValue() (any, error) {
	kind, text := token.Type, token.Text
	if kind == BackticksQuoted || kind == DoubleQuoted || kind == SingleQuoted {
		return text[1 : len(text)-1], nil
	}
	if kind == Numeric {
		if isInteger(text) {
			return strconv.ParseInt(text, 10, 64)
		} else if isDecimal(text) {
			return strconv.ParseFloat(text, 64)
		}
	}
	return text, nil
}

func determineCodePosition(inputs CharSlice, start int) (int, int) {
	lineNo := 1
	for pos := 1; pos < start; pos++ {
		r := inputs[pos]
		if r == '\n' {
			lineNo++
		}
	}
	columnNo := max(start-lineNo, 1)
	return lineNo, columnNo
}

func hasMore(inputs CharSlice, pos *int) bool {
	return *pos < len(inputs)
}

func isDecimal(text string) bool {
	return regexp.MustCompile(`^-?[0-9]+(\.[0-9]+)?$`).MatchString(text)
}

func isInteger(text string) bool {
	return regexp.MustCompile(`^-?[0-9]+$`).MatchString(text)
}

func isWhitespace(inputs CharSlice, pos *int) bool {
	chars := CharSlice{'\t', '\n', '\r', ' '}
	return hasMore(inputs, pos) && slices.Contains(chars, inputs[*pos])
}

func makeToken(inputs CharSlice, start int, end int, tokenType TokenType) *Token {
	lineNumber, columnNumber := determineCodePosition(inputs, start)
	return &Token{
		Text:         string(inputs[start:end]),
		Type:         tokenType,
		Start:        start,
		End:          end,
		LineNumber:   lineNumber,
		ColumnNumber: columnNumber,
	}
}

func nextToken(inputs CharSlice, pos *int) *Token {
	parsers := []ParserFunction{
		nextBackticksQuotedToken,
		nextDoubleQuotedToken,
		nextSingleQuotedToken,
		nextNumericToken,
		nextOperatorToken,
		nextAlphaNumericToken,
		nextSymbolToken,
	}

	// search the parsers until we find a token
	for _, parser := range parsers {
		token := parser(inputs, pos)
		if token != nil {
			return token
		}
	}
	return nil
}

func nextAlphaNumericToken(inputs CharSlice, pos *int) *Token {
	return nextEligibleToken(inputs, pos, AlphaNumeric, func(CharSlice, *int) bool {
		r := inputs[*pos]
		return unicode.IsLetter(r) || unicode.IsNumber(r)
	})
}

func nextEligibleToken(inputs CharSlice, pos *int, tokenType TokenType, isValid ValidationFunction) *Token {
	start := *pos
	if hasMore(inputs, pos) && isValid(inputs, pos) {
		for ; hasMore(inputs, pos) && isValid(inputs, pos); *pos++ {
			// nothing to do here
		}
		end := *pos
		return makeToken(inputs, start, end, tokenType)
	}
	return nil
}

func nextBackticksQuotedToken(inputs CharSlice, pos *int) *Token {
	return nextQuotedStringToken(inputs, pos, BackticksQuoted, '`')
}

func nextDoubleQuotedToken(inputs CharSlice, pos *int) *Token {
	return nextQuotedStringToken(inputs, pos, DoubleQuoted, '"')
}

func nextGlyphToken(inputs CharSlice, pos *int, tokenType TokenType, isValid ValidationFunction) *Token {
	start := *pos
	if hasMore(inputs, pos) && isValid(inputs, pos) {
		*pos++
		end := *pos
		return makeToken(inputs, start, end, tokenType)
	}
	return nil
}

func nextNumericToken(inputs CharSlice, pos *int) *Token {
	return nextEligibleToken(inputs, pos, Numeric, func(CharSlice, *int) bool {
		return unicode.IsNumber(inputs[*pos])
	})
}

func nextQuotedStringToken(inputs CharSlice, pos *int, tokenType TokenType, q rune) *Token {
	// if the current rune == `q` then capture all runes until we encounter `q` again
	if start := *pos; hasMore(inputs, pos) && inputs[start] == q {
		for *pos++; hasMore(inputs, pos) && (inputs[*pos] != q); *pos++ {
			// nothing to do here
		}
		*pos++
		end := *pos
		return makeToken(inputs, start, end, tokenType)
	}
	return nil
}

func nextOperatorToken(inputs CharSlice, pos *int) *Token {
	return nextGlyphToken(inputs, pos, Operator, func(inputs CharSlice, pos *int) bool {
		chars := CharSlice{'!', '%', '*', '&', '/', '+', '-', '<', '>', '=', '[', ']', '(', ')'}
		return slices.Contains(chars, inputs[*pos])
	})
}

func nextSingleQuotedToken(inputs CharSlice, pos *int) *Token {
	return nextQuotedStringToken(inputs, pos, SingleQuoted, '\'')
}

func nextSymbolToken(inputs CharSlice, pos *int) *Token {
	return nextGlyphToken(inputs, pos, Symbol, hasMore)
}
