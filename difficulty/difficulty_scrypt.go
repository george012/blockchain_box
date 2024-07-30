package difficulty

import (
	"encoding/hex"
	"math/big"
)

var maxTargetLTC = new(big.Int)

func init() {
	maxTargetLTC.SetString("00000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
}

// GetLTCTargetHex 将难度值转换为目标值的十六进制字符串（莱特币）
func GetLTCTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)
	target := new(big.Int).Div(maxTargetLTC, difficulty)
	return "0x" + hex.EncodeToString(target.Bytes())
}

// LTCTargetHexToDiff 将目标值的十六进制字符串转换为难度值（莱特币）
func LTCTargetHexToDiff(diffHexStr string) *big.Int {
	targetBytes, _ := hex.DecodeString(diffHexStr[2:])
	target := new(big.Int).SetBytes(targetBytes)
	return new(big.Int).Div(maxTargetLTC, target)
}
