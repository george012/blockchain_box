package coin_flags

import "github.com/btcsuite/btcd/chaincfg"

var (
	bitcoinParams = &chaincfg.MainNetParams

	liteCoinParams = &chaincfg.Params{
		PubKeyHashAddrID:        0x30, // en: P2PKH (Pay-to-PubKey-Hash) starts with L; zh-CN: P2PKH 地址以 L 开头
		ScriptHashAddrID:        0x32, // en: P2SH (Pay-to-Script-Hash) starts with M; zh-CN: P2SH 地址以 M 开头
		PrivateKeyID:            0xB0, // en: Private key starts with 6 (uncompressed) or T (compressed); zh-CN: 私钥以 6（未压缩）或 T（压缩）开头
		WitnessPubKeyHashAddrID: 0x06, // en: P2WPKH (Pay-to-Witness-PubKey-Hash) starts with ltc1; zh-CN: P2WPKH 地址以 ltc1 开头
		WitnessScriptHashAddrID: 0x0A, // en: P2WSH (Pay-to-Witness-Script-Hash) starts with ltc1; zh-CN: P2WSH 地址以 ltc1 开头
	}

	dogeCoinParams = &chaincfg.Params{
		PubKeyHashAddrID:        0x1E, // en: P2PKH starts with D; zh-CN: P2PKH 地址以 D 开头
		ScriptHashAddrID:        0x16, // en: P2SH starts with 9; zh-CN: P2SH 地址以 9 开头
		PrivateKeyID:            0x9E, // en: Private key starts with 6 (uncompressed) or Q (compressed); zh-CN: 私钥以 6（未压缩）或 Q（压缩）开头
		WitnessPubKeyHashAddrID: 0x00, // en: SegWit not supported, set to 0x00; zh-CN: 不支持 SegWit，设置为 0x00
		WitnessScriptHashAddrID: 0x00, // en: SegWit not supported, set to 0x00; zh-CN: 不支持 SegWit，设置为 0x00
	}

	bellsCoinParams = &chaincfg.Params{
		PubKeyHashAddrID:        0x19, // en: P2PKH starts with B; zh-CN: P2PKH 地址以 B 开头
		ScriptHashAddrID:        0x1E, // en: P2SH starts with C; zh-CN: P2SH 地址以 C 开头
		PrivateKeyID:            0x99, // en: Private key starts with 6 (uncompressed); zh-CN: 私钥以 6 开头
		WitnessPubKeyHashAddrID: 0x00, // en: SegWit not supported, set to 0x00; zh-CN: 不支持 SegWit，设置为 0x00
		WitnessScriptHashAddrID: 0x00, // en: SegWit not supported, set to 0x00; zh-CN: 不支持 SegWit，设置为 0x00
	}
)

// GetBIP32ChainConfig
/*
	en: get bip-32 chainConfigure;
	zh-CN: 获取BIP-32 链参数;
	@return [☑][type: chaincfg.Params] en: chainConfigure ;zh-CN: 遵循BIP-32的btcd支持的链参数;
*/
func (cf CoinFlag) GetBIP32ChainConfig() *chaincfg.Params {
	switch cf {
	case CoinFlagBTC:
		return bitcoinParams
	case CoinFlagLTC:
		return liteCoinParams
	case CoinFlagDOGE:
		return dogeCoinParams
	case CoinFlagBEL:
		return bellsCoinParams
	default:
		return bitcoinParams
	}
}
