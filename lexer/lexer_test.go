package lexer

import (
	"pomidor/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `пять = 5
десять = 10

сложение = Функция(x, y) (
  x + y
)

результат = сложение(пять, десять)
!-/*5
5 < 10 > 5

Если (5 < 10) (
	Вернуть Да
) Иначе (
	Вернуть Нет
)

10 == 10
10 != 9

"foobar"
"foo bar"
Массив(1, 2)
Пары("foo":"bar", "bar":"foo")
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "пять"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.NEW_LINE, "\n"},
		{token.IDENT, "десять"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.IDENT, "сложение"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "Функция"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.NEW_LINE, "\n"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.NEW_LINE, "\n"},
		{token.RPAREN, ")"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.IDENT, "результат"},
		{token.ASSIGN, "="},
		{token.IDENT, "сложение"},
		{token.LPAREN, "("},
		{token.IDENT, "пять"},
		{token.COMMA, ","},
		{token.IDENT, "десять"},
		{token.RPAREN, ")"},
		{token.NEW_LINE, "\n"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.NEW_LINE, "\n"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.IF, "Если"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.NEW_LINE, "\n"},
		{token.RETURN, "Вернуть"},
		{token.TRUE, "Да"},
		{token.NEW_LINE, "\n"},
		{token.RPAREN, ")"},
		{token.ELSE, "Иначе"},
		{token.LPAREN, "("},
		{token.NEW_LINE, "\n"},
		{token.RETURN, "Вернуть"},
		{token.FALSE, "Нет"},
		{token.NEW_LINE, "\n"},
		{token.RPAREN, ")"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.NEW_LINE, "\n"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.NEW_LINE, "\n"},
		{token.NEW_LINE, "\n"},
		{token.STRING, "foobar"},
		{token.NEW_LINE, "\n"},
		{token.STRING, "foo bar"},
		{token.NEW_LINE, "\n"},
		{token.ARRAY, "Массив"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.NEW_LINE, "\n"},
		{token.HASH, "Пары"},
		{token.LPAREN, "("},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.COMMA, ","},
		{token.STRING, "bar"},
		{token.COLON, ":"},
		{token.STRING, "foo"},
		{token.RPAREN, ")"},
		{token.NEW_LINE, "\n"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q, literal=%q", i, tt.expectedType, tok.Type, tt.expectedLiteral)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
