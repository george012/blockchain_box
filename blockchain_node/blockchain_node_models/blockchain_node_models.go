/*
Package blockchain_node_models blockchain Node Models
*/
package blockchain_node_models

import (
	"github.com/george012/blockchain_box/blockchain_node/blockchain_node_flag"
	"github.com/george012/blockchain_box/coin_flags"
	"github.com/george012/blockchain_box/service_info"
)

type BlockchainNodeInfo struct {
	SystemdServiceInfo service_info.ServiceInfoBase            // Systemd系统服务信息
	BlockChainNodeFlag blockchain_node_flag.BlockChainNodeFlag // 区块节点标识
	CoinFlag           coin_flags.CoinFlag                     // 币种标识
	WorkSpaceDir       string                                  // 工作目录
}

func (bcn *BlockchainNodeInfo) aTestFunc() {

}

func NewBlockChainNodeInfo(coinFlag coin_flags.CoinFlag) {

}
