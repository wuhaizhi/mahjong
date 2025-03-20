// 玩家
package biz

type Player struct {
	AccountID int

	Name string
	// 手牌
	HandCards []MahjongTile
	// 打出牌
	Discards []MahjongTile
	// 明杠
	ExposedKong []MahjongTile
	// 暗杠
	ConcealedKong []MahjongTile
	// 碰牌
	Pongs []MahjongTile
	// 吃牌
	Chows []MahjongTile
	// 是否听牌
	ReadyHand bool
	// 座位
	Location int
}
