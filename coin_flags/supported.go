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
		CoinFlagPEP.CoinName():   CoinFlagPEP,
		CoinFlagAUS.CoinName():   CoinFlagAUS,
		CoinFlagDOGM.CoinName():  CoinFlagDOGM,
		CoinFlagEAC.CoinName():   CoinFlagEAC,
		CoinFlagDINGO.CoinName(): CoinFlagDINGO,
		CoinFlagLKY.CoinName():   CoinFlagLKY,
		CoinFlagJKC.CoinName():   CoinFlagJKC,
		CoinFlagCRC.CoinName():   CoinFlagCRC,
		CoinFlagXMY.CoinName():   CoinFlagXMY,
		CoinFlagSHIC.CoinName():  CoinFlagSHIC,
		CoinFlagBRC.CoinName():   CoinFlagBRC,
		CoinFlagFLIN.CoinName():  CoinFlagFLIN,
		CoinFlagBONC.CoinName():  CoinFlagBONC,
		CoinFlagDEV.CoinName():   CoinFlagDEV,
		CoinFlagBQC.CoinName():   CoinFlagBQC,
		CoinFlagMARS.CoinName():  CoinFlagMARS,
	}
)

func IsCoinSupported(coinName string) bool {
	_, exists := SupportedCoinsMap[coinName]
	return exists
}
