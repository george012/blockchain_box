package coin_flags

// CoinRedisPrefixName 币种对应redis的前缀名称
func (cf CoinFlag) CoinRedisPrefixName() string {
	switch cf {
	case CoinFlagBTC:
		return "btc"
	case CoinFlagLTC:
		return "ltc"
	case CoinFlagDOGE:
		return "doge"
	case CoinFlagETC:
		return "etc"
	case CoinFlagETHW:
		return "ethw"
	case CoinFlagZIL:
		return "zil"
	case CoinFlagOCTA:
		return "octa"
	case CoinFlagMETA:
		return "meta"
	case CoinFlagCAU:
		return "cau"
	case CoinFlagDNX:
		return "dnx"
	case CoinFlagETH:
		return "eth"
	case CoinFlagSOL:
		return "sol"
	case CoinFlagFIL:
		return "fil"
	case CoinFlagAleo:
		return "aleo"
	case CoinFlagBEL:
		return "bel"
	case CoinFlagNone:
		return "none"
	default:
		return "none"
	}
}
