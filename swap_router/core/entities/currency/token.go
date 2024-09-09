package currency

import (
	"math/big"
	"strings"
)

type Token struct {
	Currency
	buyFeeBps  *big.Int
	sellFeeBps *big.Int
}

func NewToken(chainId int, address string, decimals int, symbol string, name string, bypassChecksum *bool, buyFeeBps *big.Int, sellFeeBps *big.Int) *Token {
	newToken := &Token{
		Currency: Currency{
			ChainId:  chainId,
			Decimals: decimals,
			Symbol:   symbol,
			Name:     name,
		},
	}

	bigZero := big.NewInt(0)
	if bypassChecksum != nil && *bypassChecksum {
		newToken.Currency.address = checkValidAddress(address)
	} else {
		newToken.Currency.address = validateAndParseAddress(address)
	}

	if buyFeeBps != nil {
		//invariant(buyFeeBps.gte(bigZero), 'NON-NEGATIVE FOT FEES');
	}
	if sellFeeBps != nil {
		//invariant(sellFeeBps.gte(bigZero), 'NON-NEGATIVE FOT FEES')
	}
	newToken.buyFeeBps = buyFeeBps
	newToken.sellFeeBps = sellFeeBps

	return newToken
}

func (t *Token) equals(other Currency) bool {
	return other.IsToken && t.ChainId == other.ChainId && strings.ToLower(t.address) == strings.ToLower(other.address)
}

// tolelom: 토큰 2개 비교해서 이름 순으로 정렬해서 pool 이름에 접근하려는 것으로 생각되는데
// 함수명이 직관적이지 않다... + 바로 함수명 리턴하는 함수 있어도 좋을 거 같은데
func (t *Token) sortsBefore(other Token) bool {
	//invariant(this.chainId === other.chainId, 'CHAIN_IDS');
	//invariant(
	//	this.address.toLowerCase() !== other.address.toLowerCase(),
	//	'ADDRESSES',
	//);
	if t.ChainId != other.ChainId {
		// other chain err
	}
	if t.address == other.address {
		// same token err
	}
	return strings.ToLower(t.address) < strings.ToLower(other.address)
}

// tolelom: 추가한 함수 sortsBefore() 대체용으로 추가
func (t *Token) SortsByLowerAddress(other Token) (Token, Token) {
	if t.ChainId != other.ChainId {
		// other chain err
	}
	if t.address == other.address {
		// same token err
	}

	if strings.ToLower(t.Currency.address) < strings.ToLower(other.Currency.address) {
		return *t, other
	}
	return other, *t
}

func (t *Token) GetToken() Token {
	return *t
}
