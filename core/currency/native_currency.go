package currency

type NativeCurrency struct {
	BaseCurrency
}

func (n *NativeCurrency) getToken() Token {
	return Token{} // 임시
}
