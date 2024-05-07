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

const (
	WalletAddressWithNone = ""
	WalletAddressWithBTC  = ""
	WalletAddressWithLTC  = ""
	WalletAddressWithDOGE = ""
	WalletAddressWithETC  = "0xf9abbad9d103522bcdc72dd38f009c6aa22f82ed"
	WalletAddressWithETHW = "0xf9abbad9d103522bcdc72dd38f009c6aa22f82ed"
	WalletAddressWithZIL  = ""
	WalletAddressWithOCTA = "0xf9abbad9d103522bcdc72dd38f009c6aa22f82ed"
	WalletAddressWithMETA = ""
	WalletAddressWithCAU  = ""
)

var WalletAddressWithZILArray = []string{
	"5be2881b75b3b942f3bff3ac7408c03104c8d3a7",
	"95eb7bcdafa549966c99910aac277358a7e91183",
	"d53c22fd0f8e6f48e85adf5a42a40bcc81ff4ca3",
	"145f877bc4dec55162d3bdb93a737279a11c9349",
	"823330e9ad4ea0efccd791d1461752fd3482abae",
	"e57edf4049eea6868112d5c37e390815390a4bc1",
	"47abd0b6ff13fe397985e68c2b8ff90418efc571",
	"e0dd79b6fe4628abb58c4f644c406780e5a0f46f",
	"2bb4243ecb30d979cc91ba8d54b1b457a08a4161",
	"138de8c4b1cba7bd1a452296cabab8a704390949",
	"b18f15091cc0337776c9028d9361c3d99f1d5732",
	"645552337f6c1ae90d11e8f94863f0d94bee4032",
	"80b641af17ec6833c2651fe83af15bf41f6f9215",
	"a32fa2ba65b88a7e3d434e764b8cde9397a2ac9f",
	"5f6a213aa8f0d591d5b83d9ddcb120161617d615",
	"a244552387d215b37e0afb7aec4ec5177f061a0e",
	"39e649a85ddd6753472fd6c8810f95ccc7f9713a",
	"da073e29a0d64d691e8ceec7a93ab7d8467e8ea7",
	"4312f0da580f6c1d258fa36e665fb5f1e21e85f7",
	"7db5021f6ed43ab4a2d285b18e31fb548236c5fd",
	"d32e104732d3573ada67bd63bb6c5982d38ec2ba",
	"16d588d55ecb3bfd3bdba3d1884a646ac2c3fd4a",
	"e7ac0335d0863663372c7677688b8200355b409c",
	"1df6d3e521254faab7712aa96f4d16d0d9b11cec",
	"b37c1919b10b579256b460ee4f6049b58d94e448",
	"173b7d505fb6056d7a3ff09f905f6f924d3f89c4",
	"36ce645dedd47d2de92d6645a01fadf82bf21efc",
	"6528bc6537d3ab5a9da4dccdab8b9836651cc112",
	"c797b0c4d586a2b00cc666980d5dbffe3dfa1819",
	"b4c7982976513409401790d7da0defb871d2ba81",
	"ea9f588f61488f0b68d1d836a738dcd6df53f41a",
	"8aa68ad7bbcd7aa9f9076a642642cff94d6296ef",
	"90feb7d90fb61fc4a2ce7ff985aa42c4255349ad",
	"8f43435a51db04391af6c3925dcc98a596a0bf26",
	"edff2ceaaa432459c241942fbe008c9531114445",
	"1643c90bb1dacda90a17ab010a7e6f7c083d9672",
	"429bc0af97e2df3c86f8e8bf973a67db02f8f974",
	"6c654a4b41dbf42a80f0a2502d9808d5ca01ea47",
	"7f50c517bd58e860f46f889391a0a0a903a08bec",
	"88a20d3f1194153fee50e4f28f65c829db3ef24c",
	"e1bbbca79a2386781e7cef1d27cd4553cc2ec711",
	"b2fc578d7e682b593296850604d1ad06508c9fd3",
	"87242152711596d60a94e20b3659c52ed7793781",
	"76fe43573b01334c1db162625c71ebaa34e961b6",
	"07517f797ffbb6d233b16abb4f03de58c77db842",
	"7eef58a58814d216e02bbb338ebd5d547c896cf1",
	"8207b22f9a0529b057ea384746ede4bf27ca1b5d",
	"c8b42d2ca53f99f93e85e1ab034680b706e1ef4e",
	"83b398f782c19cd7e0da12d616e0cd96fc730558",
	"a0364f214bc41e06b571299a7d2ccaebfb9e6a3f",
}

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

// GetWalletAddressFromNodeFlag 根据nodeFlag获取钱包地址
func (bcnf BlockChainNodeFlag) GetWalletAddressFromNodeFlag() string {
	switch bcnf {
	case BlockChainNodeFlagBTC:
		return WalletAddressWithBTC
	case BlockChainNodeFlagLTC:
		return WalletAddressWithLTC
	case BlockChainNodeFlagDOGE:
		return WalletAddressWithDOGE
	case BlockChainNodeFlagETC:
		return WalletAddressWithETC
	case BlockChainNodeFlagETHW:
		return WalletAddressWithETHW
	case BlockChainNodeFlagZIL:
		return WalletAddressWithZIL
	case BlockChainNodeFlagOCTA:
		return WalletAddressWithOCTA
	case BlockChainNodeFlagMETA:
		return WalletAddressWithMETA
	case BlockChainNodeFlagCAU:
		return WalletAddressWithCAU
	default:
		return WalletAddressWithNone
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
