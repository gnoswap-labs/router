package entities

// 원문에서는 Currency는 Token | NativeCurrency 이고
// Token과 NativeCurrency 모두 BaseCurrency를 상속받는다.
// 하지만 여기서의 Currency는 NativeCurrency와 Token에 포함된다.
type Currency interface {
	Wrapped() Token
}
