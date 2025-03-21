package biz

// 牌型
const (
	Dot           = iota + 1 // 筒牌
	Bamboo                   // 条牌
	Character                // 万牌
	East                     // 东风
	South                    // 南风
	Wast                     // 西风
	North                    // 北风
	Red                      // 红中
	Green                    // 发财
	White                    // 白板
	Spring                   // 春
	Summer                   // 夏
	Autumn                   // 秋
	Winter                   // 冬
	Plum                     // 梅
	Orchid                   // 兰
	Bamboo01                 // 竹
	Chrysanthemum            // 菊
)

// 牌型
const (
	DotName           = "筒"          // 筒牌
	BambooName        = "条"          // 条牌
	CharacterName     = "萬"          // 万牌
	EastName          = "\U0001F000" // 东风
	SouthName         = "\U0001F001" // 南风
	WastName          = "\U0001F002" // 西风
	NorthName         = "\U0001F003" // 北风
	RedName           = "\U0001F004" // 红中
	GreenName         = "\U0001F005" // 发财
	WhiteName         = "\U0001F006" // 白板
	SpringName        = "\U0001F026" // 春
	SummerName        = "\U0001F027" // 夏
	AutumnName        = "\U0001F028" // 秋
	WinterName        = "\U0001F029" // 冬
	OrchidName        = "\U0001F022" // 兰
	PlumName          = "\U0001F023" // 梅
	Bamboo01Name      = "\U0001F024" // 竹
	ChrysanthemumName = "\U0001F025" // 菊
)

// 麻将种类
const (
	// 花色 条 筒 万
	Suit = iota + 1
	// 字牌和箭牌 东 南 西 北 中 发 白
	Honor
	// 花  春夏秋冬 梅兰竹菊
	Flower
)

// 麻将各类
type MahjongKind int

// 麻将牌型
type MahjongType int

// 牌
type MahjongTile struct {
	Type   MahjongType
	Number int
	Name   string
}

var (
	SuitCards       = []MahjongType{Dot, Bamboo, Character}
	HonorCards      = []MahjongType{East, South, Wast, North, Red, Green, White}
	FlowerCards     = []MahjongType{Spring, Summer, Autumn, Winter, Plum, Orchid, Bamboo01, Chrysanthemum}
	SuitCardNames   = []string{DotName, BambooName, CharacterName}
	HonorCardNames  = []string{EastName, SouthName, WastName, NorthName, RedName, GreenName, WhiteName}
	FlowerCardNames = []string{SpringName, SummerName, AutumnName, WinterName, PlumName, OrchidName, Bamboo01Name, ChrysanthemumName}
	DotTiles        = []string{"\U0001F019", "\U0001F01A", "\U0001F01B", "\U0001F01C", "\U0001F01D", "\U0001F01E", "\U0001F01F", "\U0001F020", "\U0001F021"}
	BambooTiles     = []string{"\U0001F010", "\U0001F011", "\U0001F012", "\U0001F013", "\U0001F014", "\U0001F015", "\U0001F016", "\U0001F017", "\U0001F018"}
	CharacterTiles  = []string{"\U0001F007", "\U0001F008", "\U0001F009", "\U0001F00A", "\U0001F00B", "\U0001F00C", "\U0001F00D", "\U0001F00E", "\U0001F00F"}
)

// 规则类型 0 和牌方式 1 字牌系列 2 序数牌系列 3 刻系列 4 七对系列 5 花色组合系列 6 全带系列 8 不靠系列 9 特殊系列
type RuleType int
