package difficulty

import (
	"math/big"
)

// GetSHA256dDifficulty 根据目标值计算比特币或莱特币的难度
func GetSHA256dDifficulty(target *big.Int) *big.Float {
	maxTarget := new(big.Int)
	maxTarget.SetString("00000000FFFF0000000000000000000000000000000000000000000000000000", 16)
	diff := new(big.Float).Quo(new(big.Float).SetInt(maxTarget), new(big.Float).SetInt(target))
	return diff
}

// GetSHA256dTarget 根据难度计算目标值
func GetSHA256dTarget(difficulty *big.Int) *big.Int {
	maxTarget := new(big.Int)
	maxTarget.SetString("00000000FFFF0000000000000000000000000000000000000000000000000000", 16)
	target := new(big.Int).Div(maxTarget, difficulty)
	return target
}
