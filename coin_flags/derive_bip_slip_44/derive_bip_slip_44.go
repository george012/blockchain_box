package derive_bip_slip_44

import (
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/george012/blockchain_box/coin_flags"
)

const (
	Derive_BIP_SLIP_Tag = 44 // 默认 BIP执行标准为BIP-44 or SLIP-44
)

// DeriveBIP44PathByCoinFlag 获取 币种 硬化路径 从 coinFlag
func DeriveBIP44PathByCoinFlag(coinFlag coin_flags.CoinFlag) []uint32 {
	path := []uint32{Derive_BIP_SLIP_Tag | hdkeychain.HardenedKeyStart, coinFlag.GetBIPSLIPCoinNumber() | hdkeychain.HardenedKeyStart, 0 | hdkeychain.HardenedKeyStart, 0, 0}
	return path
}

// DeriveBIP44PathByCoinName 获取 币种 硬化路径 从 coinName
func DeriveBIP44PathByCoinName(coinName string) []uint32 {
	return DeriveBIP44PathByCoinFlag(coin_flags.GetCoinFlagByCoinName(coinName))
}
