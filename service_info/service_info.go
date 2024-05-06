package service_info

import "time"

type ServiceStatus int

const (
	ServiceStatusNone          ServiceStatus = iota
	ServiceStatusActiveRunning               //服务正在运行
	ServiceStatusActiveExited                //服务已经启动并成功退出
	ServiceStatusActiveStopped               //服务已经启动但已经停止
	ServiceStatusActiveDead                  //服务已经停止并且不会自动重启
	ServiceStatusActivating                  //服务正在启动过程中
	ServiceStatusDeactivating                //服务正在停止过程中
	ServiceStatusInactive                    //服务未运行

)

func (ss ServiceStatus) String() string {
	switch ss {
	case ServiceStatusActiveRunning:
		return "active (running)"
	case ServiceStatusActiveExited:
		return "active (exited)"
	case ServiceStatusActiveStopped:
		return "active (stopped)"
	case ServiceStatusActiveDead:
		return "active (dead)"
	case ServiceStatusActivating:
		return "activating"
	case ServiceStatusDeactivating:
		return "deactivating"
	case ServiceStatusInactive:
		return "inactive"
	default:
		return "None"
	}
}

func GetServiceStatusNumberForStr(str string) ServiceStatus {
	switch str {
	case "active (running)":
		return ServiceStatusActiveRunning
	case "active (exited)":
		return ServiceStatusActiveExited
	case "active (stopped)":
		return ServiceStatusActiveStopped
	case "active (dead)":
		return ServiceStatusActiveDead
	case "activating":
		return ServiceStatusActivating
	case "deactivating":
		return ServiceStatusDeactivating
	case "inactive":
		return ServiceStatusInactive
	default:
		return ServiceStatusNone
	}
}

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
		return "none"
	}
}

func GetServiceFlagNumberForStr(str string) ServiceFlag {
	switch str {
	case "custom_service":
		return ServiceFlagCustomService
	case "third_party_service":
		return ServiceFlagThirdParty
	default:
		return ServiceFlagNone
	}
}

type ServiceInfoBase struct {
	ServiceFlagNone   ServiceFlag   // 服务标识：自定义 or 第三方标准服务
	ServiceStatus     ServiceStatus // 服务运行状态
	ServiceStaredTime time.Time     // 启动时间
	ServiceRemark     string        // 服务备注名
	ServiceRunPath    string        // 服务程序路径
	ServiceDataPath   string        // 服务数据文件路径
}
