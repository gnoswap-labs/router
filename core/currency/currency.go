package currency

// Currency는 Token | NativeCurrency
type Currency interface {
	GetToken() Token
}
