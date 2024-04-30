package server_flags

import "time"

type ServiceFlag int

const (
	ServiceFlagNone ServiceFlag = iota // 未知服务
	ServiceFlagThirdParty
	ServiceFlagCustomService // 自定义服务
)

func (rm ServiceFlag) String() string {
	switch rm {
	case ServiceFlagCustomService:
		return "custom_service"
	case ServiceFlagThirdParty:
		return "third_party_service"
	default:
		return "None"
	}
}

type ServiceInfoBase struct {
	ServiceFlagNone   ServiceFlag // 服务标识：自定义 or 第三方标准服务
	ServiceStaredTime time.Time   // 启动时间
	ServiceRemark     string      // 服务备注名
	ServiceRunPath    string      // 服务程序路径
	ServiceDataPath   string      // 服务数据文件路径
}
