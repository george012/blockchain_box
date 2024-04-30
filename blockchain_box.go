/*
Package blockchain_box blockchain tool box(zh-cn:自定义blockchain工具箱)
*/
package blockchain_box

import "github.com/george012/blockchain_box/config"

func GetVersion() string {
	return config.ProjectVersion
}
