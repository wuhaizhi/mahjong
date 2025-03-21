package player

import "github.com/wuhaizhi/mahjong/internal/biz"

type Player struct {
	AccountID int

	Name string
	// 当前活动牌
	CurrTile biz.MahjongTile
	// 手牌
	HandCards []biz.MahjongTile
	// 打出牌
	Discards []biz.MahjongTile
	// 明杠
	ExposedKong []biz.MahjongTile
	// 暗杠
	ConcealedKong []biz.MahjongTile
	// 碰牌
	Pongs []biz.MahjongTile
	// 吃牌
	Chows []biz.MahjongTile
	// 是否听牌
	ReadyHand bool
	// 座位
	Location int
}
