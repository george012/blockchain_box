package src_flags

type SrcFlag int // dataSourceFlag来源

const (
	SrcFlagNone  SrcFlag = iota //未知
	SrcFlagOwn                  //自有
	SrcFlagThird                //第三方
)
