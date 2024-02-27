package main

import (
	"log"
	"testing"
)

func TestDetermineCodePosition(t *testing.T) {
	text := "this\n is\n 1 \n\"way of the world\"\n ! `x` + '%' $"
	start := 32
	lineNo, columnNo := DetermineCodePosition(CharSlice(text), start)
	if lineNo != 5 {
		t.Errorf("lineNo (%d) was not 5", lineNo)
	}
	if columnNo != 2 {
		t.Errorf("columnNo (%d) was not 2", columnNo)
	}
}

func TestParse(t *testing.T) {
	text := "this\n is\n 1 \n\"way of the world\"\n - `x` + '%'"
	next := Parse(text)
	for tok := next(); tok != nil; tok = next() {
		log.Printf("token: %v", AsJSON(tok))
	}
}

func TestParseFully(t *testing.T) {
	text := "this\n is\n 1 \n\"way of the world\"\n - `x` + '%'"
	tokens, err := ParseFully(text)
	if err != nil {
		t.Fatal(err.Error())
	}

	for i, tok := range tokens {
		value, err := tok.GetValue()
		if err != nil {
			t.Fatal(err.Error())
		}
		log.Printf("token[%d]: |%v| %v", i, value, AsJSON(tok))
	}
}
