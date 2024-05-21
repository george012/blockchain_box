package cert_info

import "time"

// Certificate 定义自定义服务器证书的结构体
type Certificate struct {
	DomainName          string    // 域名名称
	ValidFrom           time.Time // 证书有效期开始时间
	ValidUntil          time.Time // 证书有效期截止时间
	DaysUntilExpiry     int       // 证书剩余有效天数
	CertificateFilePath string    // 证书文件路径
}
