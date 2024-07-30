package difficulty

import (
	"encoding/hex"
	"math/big"
)

var maxTargetSha256d = new(big.Int)

func init() {
	maxTargetSha256d.SetString("00000000FFFF0000000000000000000000000000000000000000000000000000", 16)
}

// GetTargetHexBySha256d 将难度值转换为目标值的十六进制字符串（比特币 Sha256d 算法）
func GetTargetHexBySha256d(diff int64) string {
	difficulty := big.NewInt(diff)
	target := new(big.Int).Div(maxTargetSha256d, difficulty)
	return "0x" + hex.EncodeToString(target.Bytes())
}

// TargetHexToDiffBySha256d 将目标值的十六进制字符串转换为难度值（比特币 Sha256d 算法）
func TargetHexToDiffBySha256d(diffHexStr string) *big.Int {
	targetBytes, _ := hex.DecodeString(diffHexStr[2:])
	target := new(big.Int).SetBytes(targetBytes)
	return new(big.Int).Div(maxTargetSha256d, target)
}
