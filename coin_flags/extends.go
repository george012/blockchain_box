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
	case CoinFlagPEP:
		return "pep"
	case CoinFlagAUS:
		return "aus"
	case CoinFlagDOGM:
		return "dogm"
	case CoinFlagEAC:
		return "eac"
	case CoinFlagDINGO:
		return "dingo"
	case CoinFlagLKY:
		return "lky"
	case CoinFlagJKC:
		return "jkc"
	case CoinFlagCRC:
		return "crc"
	case CoinFlagXMY:
		return "xmy"
	case CoinFlagSHIC:
		return "shic"
	case CoinFlagBRC:
		return "brc"
	case CoinFlagFLIN:
		return "flin"
	case CoinFlagBONC:
		return "bonc"
	case CoinFlagDEV:
		return "dev"
	case CoinFlagBQC:
		return "bqc"
	case CoinFlagNone:
		return "none"
	default:
		return "none"
	}
}
