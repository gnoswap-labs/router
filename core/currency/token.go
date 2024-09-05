package currency

import "strings"

type Token struct {
	BaseCurrency
	Address string
}

func (t *Token) getToken() Token {
	return *t
}

// tolelom: 토큰 2개 비교해서 이름 순으로 정렬해서 pool 이름에 접근하려는 것으로 생각되는데
// 함수명이 직관적이지 않다... + 바로 함수명 리턴하는 함수 있어도 좋을 거 같은데
func (t *Token) sortsBefore(other Token) bool {
	if t.ChainId != other.ChainId {
		// other chain err
	}
	if t.address == other.address {
		// same token err
	}

	return strings.ToLower(t.Address) < strings.ToLower(other.Address)
}
