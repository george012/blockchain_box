package coin_info

var SupportCoins = []*CoinInfo{}

type CoinInfo struct {
	CoinName              string   `json:"coin_name"`
	Algorithm             string   `json:"algorithm"`
	SupportMultiAlgorithm bool     `json:"support_multi_algorithm"`
	MultiAlgorithms       []string `json:"multi_algorithms"`
	ChainID               string   `json:"chain_id"`
	ChainName             string   `json:"chain_name"`
	CoinFullName          string   `json:"coin_full_name"`
	SupportAuxMergeMining bool     `json:"support_aux_merge_mining"`
}

func (ci *CoinInfo) addSupported() {
	SupportCoins = append(SupportCoins, ci)
}

func init() {

}
