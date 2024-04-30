package server_flags

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
