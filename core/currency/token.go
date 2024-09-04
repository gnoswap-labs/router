package currency

type Token struct {
	BaseCurrency
	address string
}

func (t *Token) getToken() Token {
	return *t
}
