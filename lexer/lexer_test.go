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
);

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
"Строка";
"Ещё строка";
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
		{token.STRING, "Строка"},
		{token.SEMICOLON, ";"},
		{token.STRING, "Ещё строка"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
