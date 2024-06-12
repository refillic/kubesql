package sqlparser

import (
	"testing"
)

func TestSingleSelect(t *testing.T) {

	testString := "SELECT"

	expectedTokens := []Token {
		Token {
			Type: Select,
			Value: "SELECT",
		},
	}

	lexer := NewLexer(testString)
	actualTokens := lexer.Tokenize()

	numberOfLexedTokens := len(actualTokens)
	numberOfExpectedTokens := len(expectedTokens)

	t.Logf("Actual tokens %v", actualTokens)

	if (numberOfLexedTokens != numberOfExpectedTokens) {
		t.Errorf("Number of expected tokens (%d) doesn't match number of actual tokens (%d)", numberOfExpectedTokens, numberOfLexedTokens)
	}

	for i, actualToken := range actualTokens {
		tokenMatchesExpected := actualToken.Type == expectedTokens[i].Type && actualToken.Value == expectedTokens[i].Value
		if !tokenMatchesExpected {
			t.Errorf("Lexed token (%v) doesn't match expected token (%v)", actualToken, expectedTokens[i])
		}
	}
}

func TestSingleSelectWithOneColumn(t *testing.T) {

	testString := "SELECT a"

	expectedTokens := []Token {
		Token {
			Type: Select,
			Value: "SELECT",
		},
		Token {
			Type: Identifier,
			Value: "a",
		},
	}

	lexer := NewLexer(testString)
	actualTokens := lexer.Tokenize()

	numberOfLexedTokens := len(actualTokens)
	numberOfExpectedTokens := len(expectedTokens)

	t.Logf("Actual tokens %v", actualTokens)

	if (numberOfLexedTokens != numberOfExpectedTokens) {
		t.Errorf("Number of expected tokens (%d) doesn't match number of actual tokens (%d)", numberOfExpectedTokens, numberOfLexedTokens)
	}

	for i, actualToken := range actualTokens {
		tokenMatchesExpected := actualToken.Type == expectedTokens[i].Type && actualToken.Value == expectedTokens[i].Value
		if !tokenMatchesExpected {
			t.Errorf("Lexed token (%v) doesn't match expected token (%v)", actualToken, expectedTokens[i])
		}
	}
}

func TestSingleSelectWithMultipleColumns(t *testing.T) {

	testString := "SELECT a, b, c"

	expectedTokens := []Token {
		Token {
			Type: Select,
			Value: "SELECT",
		},
		Token {
			Type: Identifier,
			Value: "a",
		},
		Token {
			Type: Comma,
			Value: ",",
		},
		Token {
			Type: Identifier,
			Value: "b",
		},
		Token {
			Type: Comma,
			Value: ",",
		},
		Token {
			Type: Identifier,
			Value: "c",
		},
	}

	lexer := NewLexer(testString)
	actualTokens := lexer.Tokenize()

	numberOfLexedTokens := len(actualTokens)
	numberOfExpectedTokens := len(expectedTokens)

	t.Logf("Actual tokens %v", actualTokens)

	if (numberOfLexedTokens != numberOfExpectedTokens) {
		t.Errorf("Number of expected tokens (%d) doesn't match number of actual tokens (%d)", numberOfExpectedTokens, numberOfLexedTokens)
	}

	for i, actualToken := range actualTokens {
		tokenMatchesExpected := actualToken.Type == expectedTokens[i].Type && actualToken.Value == expectedTokens[i].Value
		if !tokenMatchesExpected {
			t.Errorf("Lexed token (%v) doesn't match expected token (%v)", actualToken, expectedTokens[i])
		}
	}
}

func TestSingleSelectWithFrom(t *testing.T) {

	testString := "SELECT a, b, c FROM pods"

	expectedTokens := []Token {
		Token {
			Type: Select,
			Value: "SELECT",
		},
		Token {
			Type: Identifier,
			Value: "a",
		},
		Token {
			Type: Comma,
			Value: ",",
		},
		Token {
			Type: Identifier,
			Value: "b",
		},
		Token {
			Type: Comma,
			Value: ",",
		},
		Token {
			Type: Identifier,
			Value: "c",
		},
		Token {
			Type: From,
			Value: "FROM",
		},
		Token {
			Type: Identifier,
			Value: "pods",
		},
	}

	lexer := NewLexer(testString)
	actualTokens := lexer.Tokenize()

	numberOfLexedTokens := len(actualTokens)
	numberOfExpectedTokens := len(expectedTokens)

	t.Logf("Actual tokens %v", actualTokens)

	if (numberOfLexedTokens != numberOfExpectedTokens) {
		t.Errorf("Number of expected tokens (%d) doesn't match number of actual tokens (%d)", numberOfExpectedTokens, numberOfLexedTokens)
	}

	for i, actualToken := range actualTokens {
		tokenMatchesExpected := actualToken.Type == expectedTokens[i].Type && actualToken.Value == expectedTokens[i].Value
		if !tokenMatchesExpected {
			t.Errorf("Lexed token (%v) doesn't match expected token (%v)", actualToken, expectedTokens[i])
		}
	}
}