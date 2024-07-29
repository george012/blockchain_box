package difficulty

// DiffUnit 格式显示
type DiffUnit string

const (
	DiffUnitNone  DiffUnit = "none"  // default none 未知,默认:无单位或基础单位
	DiffUnitKilo  DiffUnit = "Kilo"  // Kilo 千
	DiffUnitMega  DiffUnit = "Mega"  // Mega 兆
	DiffUnitGiga  DiffUnit = "Giga"  // Giga 吉
	DiffUnitTera  DiffUnit = "Tera"  // Tera 太
	DiffUnitPeta  DiffUnit = "Peta"  // Peta 拍
	DiffUnitExa   DiffUnit = "Exa"   // Exa 艾
	DiffUnitZetta DiffUnit = "Zetta" // Zetta 泽
	DiffUnitYotta DiffUnit = "Yotta" // Yotta 尧
)

// Shortname 获取简写
func (aType DiffUnit) Shortname() string {
	switch aType {
	case DiffUnitKilo:
		return "K"
	case DiffUnitMega:
		return "M"
	case DiffUnitGiga:
		return "G"
	case DiffUnitTera:
		return "T"
	case DiffUnitPeta:
		return "P"
	case DiffUnitExa:
		return "E"
	case DiffUnitZetta:
		return "Z"
	case DiffUnitYotta:
		return "Y"
	case DiffUnitNone:
		return ""
	default:
		return ""
	}
}

type DiffUnitFormatType int

const (
	DiffUnitFormatType1000 DiffUnitFormatType = 1000
	DiffUnitFormatType1024 DiffUnitFormatType = 1024
)
