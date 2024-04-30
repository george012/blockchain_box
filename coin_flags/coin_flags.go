/*
Package coin_flags coin flags(zh-cn:币种标识)
*/
package coin_flags

var (
	SupportedCoinsMap = map[string]CoinFlag{
		CoinFlagBTC.CoinName():  CoinFlagBTC,
		CoinFlagETC.CoinName():  CoinFlagETC,
		CoinFlagETHW.CoinName(): CoinFlagETHW,
		CoinFlagZIL.CoinName():  CoinFlagZIL,
		CoinFlagOCTA.CoinName(): CoinFlagOCTA,
		CoinFlagLTC.CoinName():  CoinFlagLTC,
		CoinFlagDOGE.CoinName(): CoinFlagDOGE,
		CoinFlagMETA.CoinName(): CoinFlagMETA,
		CoinFlagCAU.CoinName():  CoinFlagCAU,
	}
)

type CoinFlag int

const (
	CoinFlagNone CoinFlag = iota
	CoinFlagBTC
	CoinFlagLTC
	CoinFlagDOGE
	CoinFlagETC
	CoinFlagETHW
	CoinFlagZIL
	CoinFlagOCTA
	CoinFlagMETA
	CoinFlagCAU
)

func GetCoinFlagByCoinName(name string) CoinFlag {
	switch name {
	case "BTC", "BitCoin", "Bitcoin":
		return CoinFlagBTC
	case "LTC", "LiteCoin", "Litecoin":
		return CoinFlagLTC
	case "DOGE", "DogeCoin", "Dogecoin":
		return CoinFlagDOGE
	case "ETC", "Ethereum Classic", "EthereumClassic":
		return CoinFlagETC
	case "ETHW", "EthereumPoW":
		return CoinFlagETHW
	case "ZIL", "Zilliqa":
		return CoinFlagZIL
	case "OCTA", "OctaSpace":
		return CoinFlagOCTA
	case "META", "MetaChain":
		return CoinFlagMETA
	case "CAU", "Canxium":
		return CoinFlagCAU
	default:
		return CoinFlagNone
	}
}

func (cf CoinFlag) CoinName() string {
	switch cf {
	case CoinFlagBTC:
		return "BTC"
	case CoinFlagLTC:
		return "LTC"
	case CoinFlagDOGE:
		return "DOGE"
	case CoinFlagETC:
		return "ETC"
	case CoinFlagETHW:
		return "ETHW"
	case CoinFlagZIL:
		return "ZIL"
	case CoinFlagOCTA:
		return "OCTA"
	case CoinFlagMETA:
		return "META"
	case CoinFlagCAU:
		return "CAU"
	default:
		return "none"
	}
}

func (cf CoinFlag) CoinFullName() string {
	switch cf {
	case CoinFlagBTC:
		return "Bitcoin"
	case CoinFlagLTC:
		return "Litecoin"
	case CoinFlagDOGE:
		return "Dogecoin"
	case CoinFlagETC:
		return "EthereumClassic"
	case CoinFlagETHW:
		return "EthereumPoW"
	case CoinFlagZIL:
		return "Zilliqa"
	case CoinFlagOCTA:
		return "OctaSpace"
	case CoinFlagMETA:
		return "MetaChain"
	case CoinFlagCAU:
		return "Canxium"
	default:
		return "none"
	}
}

func IsCoinSupported(coinName string) bool {
	_, exists := SupportedCoinsMap[coinName]
	return exists
}