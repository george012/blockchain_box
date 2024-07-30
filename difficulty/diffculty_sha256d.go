package difficulty

import (
	"encoding/hex"
	"math/big"
)

var maxTargetBTC = new(big.Int)

func init() {
	maxTargetBTC.SetString("00000000FFFF0000000000000000000000000000000000000000000000000000", 16)
}

// GetBTCTargetHex 将难度值转换为目标值的十六进制字符串（比特币）
func GetBTCTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)
	target := new(big.Int).Div(maxTargetBTC, difficulty)
	return "0x" + hex.EncodeToString(target.Bytes())
}

// BTCTargetHexToDiff 将目标值的十六进制字符串转换为难度值（比特币）
func BTCTargetHexToDiff(diffHexStr string) *big.Int {
	targetBytes, _ := hex.DecodeString(diffHexStr[2:])
	target := new(big.Int).SetBytes(targetBytes)
	return new(big.Int).Div(maxTargetBTC, target)
}
