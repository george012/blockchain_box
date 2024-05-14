package sys_infos

type NetworkInterfaceInfo struct {
	LocalIP string // 192.168.1.123
	PubIP   string // 134.164.148.44
}

type DiskInfo struct {
	TotalSize uint64
	UsedSize  uint64
	FreeSize  uint64
}

type MemoryInfo struct {
	TotalSize uint64
	UsedSize  uint64
	FreeSize  uint64
}

type CPUInfo struct {
	ModelName string // 型号
	BrandName string // 厂商
}
