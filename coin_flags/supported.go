package coin_flags

var (
	SupportedCoinsMap = map[string]CoinFlag{
		CoinFlagBTC.CoinName():   CoinFlagBTC,
		CoinFlagETC.CoinName():   CoinFlagETC,
		CoinFlagETHW.CoinName():  CoinFlagETHW,
		CoinFlagZIL.CoinName():   CoinFlagZIL,
		CoinFlagOCTA.CoinName():  CoinFlagOCTA,
		CoinFlagLTC.CoinName():   CoinFlagLTC,
		CoinFlagDOGE.CoinName():  CoinFlagDOGE,
		CoinFlagMETA.CoinName():  CoinFlagMETA,
		CoinFlagCAU.CoinName():   CoinFlagCAU,
		CoinFlagBEL.CoinName():   CoinFlagBEL,
		CoinFlagPEPE.CoinName():  CoinFlagPEPE,
		CoinFlagAUS.CoinName():   CoinFlagAUS,
		CoinFlagDOGM.CoinName():  CoinFlagDOGM,
		CoinFlagEAC.CoinName():   CoinFlagEAC,
		CoinFlagDINGO.CoinName(): CoinFlagDINGO,
	}
)

func IsCoinSupported(coinName string) bool {
	_, exists := SupportedCoinsMap[coinName]
	return exists
}
