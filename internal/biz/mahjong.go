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
	DotName           = "筒" // 筒牌
	BambooName        = "条" // 条牌
	CharacterName     = "萬" // 万牌
	EastName          = "東" // 东风
	SouthName         = "南" // 南风
	WastName          = "西" // 西风
	NorthName         = "北" // 北风
	RedName           = "中" // 红中
	GreenName         = "發" // 发财
	WhiteName         = "白" // 白板
	SpringName        = "春" // 春
	SummerName        = "夏" // 夏
	AutumnName        = "秋" // 秋
	WinterName        = "冬" // 冬
	PlumName          = "梅" // 梅
	OrchidName        = "蘭" // 兰
	Bamboo01Name      = "竹" // 竹
	ChrysanthemumName = "菊" // 菊
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
)

// 麻将
type Mahjong struct {
	Kind []MahjongKind
}
