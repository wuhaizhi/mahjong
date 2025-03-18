package biz

// 牌型
const (
	Dot    = iota + 1
	Bamboo // 条
	Character
	East
	South
	Wast
	North
	Red
	Green
	White
	Spring
	Summer
	Autumn
	Winter
	Plum
	Orchid
	Bamboo01
	Chrysanthemum
)

// 麻将种类
const (
	// 花色 条 筒 万
	Suit = iota + 1
	// 字牌 东 南 西 北 中 发 白
	Honor
	// 花  春夏秋冬 梅兰竹菊
	Flower
)

// 麻将各类
type MahjongKind int

// 麻将牌型
type MahjongType int

// 牌
type MahjongCard struct {
	Type   MahjongType
	Number int
	Name   string
}

// 麻将
type Mahjong struct {
	Kind []MahjongKind
}

// 洗牌
func (m *Mahjong) shuffle() {
}

// 抓牌
func (m *Mahjong) Draw() {
}

// 出牌
func (m *Mahjong) Produce() {
}
