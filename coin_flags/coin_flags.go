/*
Package coin_flags coin flags(zh-cn:币种标识)
*/
package coin_flags

import (
	"fmt"
	"strings"
)

// CoinFlag
/*
	en: Blockchain currency identifier;
	zh-CN: 区块链币种标识;
*/
type CoinFlag int

const (
	CoinFlagNone CoinFlag = iota // en: Unknown Type | zh-CN: 未知类型;
	CoinFlagMETA                 // en: META | zh-CN: META;
	CoinFlagCAU                  // en: CAU | zh-CN: CAU;
	CoinFlagETC                  // en: ETC | zh-CN: ETC;
	CoinFlagZIL                  // en: ZIL | zh-CN: ZIL;
	CoinFlagETHW                 // en: ETHW | zh-CN: ETHW;
	CoinFlagBTC                  // en: BTC | zh-CN: BTC;
	CoinFlagLTC                  // en: LTC | zh-CN: LTC;
	CoinFlagDNX                  // en: DNX | zh-CN: DNX;
	CoinFlagOCTA                 // en: OCTA | zh-CN: OCTA;
	CoinFlagDOGE                 // en: DOGE | zh-CN: DOGE;
	CoinFlagETH                  // en: ETH | zh-CN: ETH;
	CoinFlagSOL                  // en: SOL | zh-CN: SOL;
	CoinFlagFIL                  // en: FIL | zh-CN: FIL;
)

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
	case CoinFlagDNX:
		return "DNX"
	case CoinFlagETH:
		return "ETH"
	case CoinFlagSOL:
		return "SOL"
	case CoinFlagFIL:
		return "FIL"
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
	case CoinFlagDNX:
		return "Dynexcoin"
	case CoinFlagETH:
		return "Ethereum"
	case CoinFlagSOL:
		return "Solana"
	case CoinFlagFIL:
		return "FileCoin"
	default:
		return "none"
	}
}

// GetBlockNodeBinaryName
/*
	en: Get blockchain node binary program name;
	zh-CN: 获取区块链节点 二进制程序 名称;
*/
func (cf CoinFlag) GetBlockNodeBinaryName() string {
	switch cf {
	case CoinFlagBTC:
		return "btcd"
	case CoinFlagLTC:
		return "litecoind"
	case CoinFlagDOGE:
		return "dogecoind"
	case CoinFlagETC:
		return "geth"
	case CoinFlagETHW:
		return "geth"
	case CoinFlagZIL:
		return "zilliqa"
	case CoinFlagOCTA:
		return "geth"
	case CoinFlagMETA:
		return "geth"
	case CoinFlagCAU:
		return "canxium"
	case CoinFlagDNX:
		return "geth"
	case CoinFlagETH:
		return "geth"
	case CoinFlagSOL:
		return "solana"
	case CoinFlagFIL:
		return "lotus"
	default:
		return "none"
	}
}

// GetBlockNodeBinarySystemdServiceName
/*
	en: Get the Systemd service name of the blockchain node binary program;
	zh-CN: 获取区块链节点 二进制程序的Systemd服务 名称;
*/
func (cf CoinFlag) GetBlockNodeBinarySystemdServiceName() string {
	sName := fmt.Sprintf("%s-%s", strings.ToLower(cf.CoinName()), cf.GetBlockNodeBinaryName())
	// TODO 对ETC做单独处理
	if cf == CoinFlagETC {
		sName = fmt.Sprintf("%s-%s", "core", cf.GetBlockNodeBinaryName())
	}
	return sName
}
