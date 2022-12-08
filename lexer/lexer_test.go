package lexer

import (
	"pomidor/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `пять = 5;
десять = 10;

сложение = Функция(x, y) (
  x + y;
)

вычитание = Функция(x, y) (Вернуть x - y);

результат = сложение(пять, десять);
!-/*5;
5 < 10 > 5;

Если (5 < 10) (
	Вернуть Да;
) Иначе (
	Вернуть Нет;
)

10 == 10;
10 != 9;

"foobar";
"foo bar";
Массив(1, 2):(1);
Пары("foo":"bar", "bar":"foo");
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "пять"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "десять"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "сложение"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "Функция"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RPAREN, ")"},
		{token.IDENT, "вычитание"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "Функция"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.RETURN, "Вернуть"},
		{token.IDENT, "x"},
		{token.MINUS, "-"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "результат"},
		{token.ASSIGN, "="},
		{token.IDENT, "сложение"},
		{token.LPAREN, "("},
		{token.IDENT, "пять"},
		{token.COMMA, ","},
		{token.IDENT, "десять"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "Если"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LPAREN, "("},
		{token.RETURN, "Вернуть"},
		{token.TRUE, "Да"},
		{token.SEMICOLON, ";"},
		{token.RPAREN, ")"},
		{token.ELSE, "Иначе"},
		{token.LPAREN, "("},
		{token.RETURN, "Вернуть"},
		{token.FALSE, "Нет"},
		{token.SEMICOLON, ";"},
		{token.RPAREN, ")"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foo bar"},
		{token.SEMICOLON, ";"},
		{token.ARRAY, "Массив"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.COLON, ":"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
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
		{token.SEMICOLON, ";"},
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
