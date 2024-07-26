package coin_flags

func (cf CoinFlag) Algorithm() string {
	switch cf {
	case CoinFlagBTC:
		return "SHA256d"
	case CoinFlagLTC:
		return "Scrypt"
	case CoinFlagDOGE:
		return "Scrypt"
	case CoinFlagETC:
		return "ETCHash"
	default:
		return "none"
	}
}
