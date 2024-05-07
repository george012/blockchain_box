/*
Package blockchain_node_flag blockchain node flags(zh-cn:区块链节点标识)
*/
package blockchain_node_flag

var (
	SupportedBlockChainNodes = map[string]BlockChainNodeFlag{
		BlockChainNodeFlagBTC.NodeServiceName():  BlockChainNodeFlagBTC,
		BlockChainNodeFlagLTC.NodeServiceName():  BlockChainNodeFlagLTC,
		BlockChainNodeFlagDOGE.NodeServiceName(): BlockChainNodeFlagDOGE,
		BlockChainNodeFlagETC.NodeServiceName():  BlockChainNodeFlagETC,
		BlockChainNodeFlagETHW.NodeServiceName(): BlockChainNodeFlagETHW,
		BlockChainNodeFlagZIL.NodeServiceName():  BlockChainNodeFlagZIL,
		BlockChainNodeFlagOCTA.NodeServiceName(): BlockChainNodeFlagOCTA,
		BlockChainNodeFlagMETA.NodeServiceName(): BlockChainNodeFlagMETA,
		BlockChainNodeFlagCAU.NodeServiceName():  BlockChainNodeFlagCAU,
	}
)

type BlockChainNodeFlag int

// server常量
const (
	BlockChainNodeFlagNone BlockChainNodeFlag = iota
	BlockChainNodeFlagBTC
	BlockChainNodeFlagLTC
	BlockChainNodeFlagDOGE
	BlockChainNodeFlagETC
	BlockChainNodeFlagETHW
	BlockChainNodeFlagZIL
	BlockChainNodeFlagOCTA
	BlockChainNodeFlagMETA
	BlockChainNodeFlagCAU
)

// GetBlockChainNodeFlagFromCoinName 根据coinName 获取节点标识
func GetBlockChainNodeFlagFromCoinName(coinName string) BlockChainNodeFlag {
	switch coinName {
	case "BTC", "BitCoin":
		return BlockChainNodeFlagBTC
	case "LTC", "LiteCoin":
		return BlockChainNodeFlagLTC
	case "DOGE", "Doge coin", "Dogecoin":
		return BlockChainNodeFlagDOGE
	case "ETC", "Ethereum Classic", "EthereumClassic":
		return BlockChainNodeFlagETC
	case "ETHW", "Ethereum PoW", "EthereumPoW":
		return BlockChainNodeFlagETHW
	case "ZIL", "Zilliqa":
		return BlockChainNodeFlagZIL
	case "OCTA", "OctaSpace":
		return BlockChainNodeFlagOCTA
	case "META", "MetaChain":
		return BlockChainNodeFlagMETA
	case "CAU", "Canxium":
		return BlockChainNodeFlagCAU
	default:
		return BlockChainNodeFlagNone
	}
}

// TODO  未定稿
// NodeServiceName 节点服务名称
func (bcnf BlockChainNodeFlag) NodeServiceName() string {
	switch bcnf {
	case BlockChainNodeFlagBTC:
		return "btcd"
	case BlockChainNodeFlagLTC:
		return "litecoind"
	case BlockChainNodeFlagDOGE:
		return "dogecoind"
	case BlockChainNodeFlagETC:
		return "core_geth"
	case BlockChainNodeFlagETHW:
		return "ethw_geth"
	case BlockChainNodeFlagZIL:
		return "zilliqa"
	case BlockChainNodeFlagOCTA:
		return "octa_geth"
	case BlockChainNodeFlagMETA:
		return "meta_geth"
	case BlockChainNodeFlagCAU:
		return "cau_geth"
	default:
		return "none"
	}
}

// GetBlockChainNodeFlagFromServiceName 根据NodeServiceName 获取节点标识
func GetBlockChainNodeFlagFromServiceName(coinName string) BlockChainNodeFlag {
	switch coinName {
	case "btcd":
		return BlockChainNodeFlagBTC
	case "litecoind":
		return BlockChainNodeFlagLTC
	case "dogecoind":
		return BlockChainNodeFlagDOGE
	case "core_geth":
		return BlockChainNodeFlagETC
	case "ethw_geth":
		return BlockChainNodeFlagETHW
	case "zilliqa":
		return BlockChainNodeFlagZIL
	case "octa_geth":
		return BlockChainNodeFlagOCTA
	case "meta_geth":
		return BlockChainNodeFlagMETA
	case "cau_geth":
		return BlockChainNodeFlagCAU
	default:
		return BlockChainNodeFlagNone
	}
}

// TODO  END 未定稿

func IsCustomService(serviceName string) bool {
	_, exists := SupportedBlockChainNodes[serviceName]
	return exists
}
