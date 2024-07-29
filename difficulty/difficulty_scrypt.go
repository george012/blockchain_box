package difficulty

import "math/big"

// GetScryptDifficulty 根据目标值计算比特币或莱特币的难度
func GetScryptDifficulty(target *big.Int) *big.Float {
	return GetSHA256dDifficulty(target)
}

// GetScryptTarget 根据难度计算目标值
func GetScryptTarget(difficulty *big.Int) *big.Int {
	return GetSHA256dTarget(difficulty)
}
