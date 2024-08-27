package entities

type Token struct {
	BaseCurrency
}

func NewToken() *Token {
	return &Token{}
}

func (t Token) Wrapped() Token {
	return t
}
