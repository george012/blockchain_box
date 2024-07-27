package coin_flags

// PowAlgorithm
/*
	en: get pow Algorithm;
	zh-CN: 获取coin 对应的 Pow算法;
	@return [☑][type: string] en: Algorithm string ;zh-CN: 算法字符串;
*/
func (cf CoinFlag) PowAlgorithm() string {
	switch cf {
	case CoinFlagBTC:
		return "SHA256d"
	case CoinFlagLTC:
		return "Scrypt"
	case CoinFlagDOGE:
		return "Scrypt"
	case CoinFlagETC:
		return "EtcHash"
	case CoinFlagETHW:
		return "Ethash"
	case CoinFlagZIL:
		return "Ethash"
	case CoinFlagOCTA:
		return "Ethash"
	case CoinFlagMETA:
		return "Ethash"
	case CoinFlagCAU:
		return "Ethash"
	case CoinFlagDNX:
		return "DynexSolve"
	case CoinFlagETH:
		return "none"
	case CoinFlagSOL:
		return "none"
	case CoinFlagFIL:
		return "none"
	case CoinFlagAleo:
		return "zkSNARK"
	default:
		return "none"
	}
}
