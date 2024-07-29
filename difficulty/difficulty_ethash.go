package difficulty

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

var pow256Ethash = math.BigPow(2, 256)

// GetTargetHex 将难度值转换为目标值的十六进制字符串
func GetTargetHex(diff int64) string {
	difficulty := big.NewInt(diff)
	diff1 := new(big.Int).Div(pow256Ethash, difficulty)
	return hexutil.Encode(diff1.Bytes())
}

// TargetHexToDiff 将目标值的十六进制字符串转换为难度值
func TargetHexToDiff(diffHexStr string) *big.Int {
	targetBytes := common.FromHex(diffHexStr)
	return new(big.Int).Div(pow256Ethash, new(big.Int).SetBytes(targetBytes))
}
