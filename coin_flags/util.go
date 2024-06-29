package coin_flags

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
	case "DNX", "Dynexcoin":
		return CoinFlagDNX
	case "ETH", "Ethereum":
		return CoinFlagETH
	default:
		return CoinFlagNone
	}
}
