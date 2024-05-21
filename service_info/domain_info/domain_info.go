package domain_info

import "github.com/george012/blockchain_box/service_info/cert_info"

// DomainInfo 域名信息
type DomainInfo struct {
	DomainName  string                            // 域名名称
	Certificate map[string]*cert_info.Certificate // 证书信息
}
