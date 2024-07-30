package difficulty

import (
	"encoding/hex"
	"math/big"
)

var maxTargetScrypt = new(big.Int)

func init() {
	maxTargetScrypt.SetString("00000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
}

// GetTargetHexByScrypt 将难度值转换为目标值的十六进制字符串（莱特币Scrypt 算法类）
func GetTargetHexByScrypt(diff int64) string {
	difficulty := big.NewInt(diff)
	target := new(big.Int).Div(maxTargetScrypt, difficulty)
	return "0x" + hex.EncodeToString(target.Bytes())
}

// TargetHexToDiffByScrypt 将目标值的十六进制字符串转换为难度值（莱特币Scrypt 算法类）
func TargetHexToDiffByScrypt(diffHexStr string) *big.Int {
	targetBytes, _ := hex.DecodeString(diffHexStr[2:])
	target := new(big.Int).SetBytes(targetBytes)
	return new(big.Int).Div(maxTargetScrypt, target)
}
