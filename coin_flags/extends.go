package coin_flags

import "strings"

// CoinRedisPrefixName 币种对应redis的前缀名称
func (cf CoinFlag) CoinRedisPrefixName() string {
	return strings.ToLower(cf.CoinName())
}
