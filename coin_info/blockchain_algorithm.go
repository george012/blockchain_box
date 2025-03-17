package coin_info

type Algorithm string

const (
	AlgorithmSHA256D Algorithm = "SHA256D"
	AlgorithmScrypt  Algorithm = "Scrypt"
	AlgorithmETHash  Algorithm = "ETHash"
	AlgorithmETCHash Algorithm = "ETCHash"
	AlgorithmRamdomX Algorithm = "RamdomX"
)
