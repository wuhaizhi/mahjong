package biz

// 麻将逻辑
//
// 条 饼  万  东西南北中发白  春夏秋冬
//
// 条  1-9  4
// 饼  1-9  4
// 万  1-9  4
//
//
//
// 顺子   ABC
// 刻子   AAA
// 对子   DD
//
//
// 和牌 n*AAA + m*ABC + s*DD
// n + m = 4, s = 1, 即可和牌。此时m和n不同时为0如果m和n同时为0， 则s必须是7，才能和牌。
// 此为基本公式，忽略特殊情况
//
//
//
// 麻将过程：
// 1、洗牌
// 2、砌牌
// 3、抓牌
// 4、打牌
//
//
// 碰吃杠
//
// 碰  三家皆可收
// 吃  只能吃上家
// 杠  三家皆可收
//
// 听牌   只差一张和牌

const (
	Dot    = iota
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

type MahjongCard struct {
	Type   int // 0 条 1 饼 2 万 3 东 4 西 5 南 6 北 7 中 8 发 9 白 10 春 11 夏 12 秋 13冬
	Number int
	Name   string
}

// 麻将
type Mahjong struct{}

// 洗牌
func (m *Mahjong) shuffle() {
}

// 抓牌
func (m *Mahjong) Draw() {
}

// 出牌
func (m *Mahjong) Produce() {
}
