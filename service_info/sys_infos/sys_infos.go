package sys_infos

type NetworkInterfaceInfo struct {
	InterfaceName string // 网卡名
	LocalIP       string // 192.168.1.123
	PubIP         string // 134.164.148.44
}

type DiskInfo struct {
	TotalSize  uint64
	UsedSize   uint64
	FreeSize   uint64
	Percentage float64 // 使用百分比
}

type MemoryInfo struct {
	TotalSize  uint64
	UsedSize   uint64
	FreeSize   uint64
	Percentage float64 // 使用百分比 举例 50% 即写成 0.5
}

type CPUInfo struct {
	ModelName  string  // 型号
	BrandName  string  // 厂商
	Percentage float64 // 使用百分比 举例 50% 即写成 0.5
	Cores      uint64  // 核心数
	Threads    uint64  //  线程数
}
