package biz

// 牌墙
type Wall struct {
	EastWall  []MahjongTile
	SouthWall []MahjongTile
	WestWall  []MahjongTile
	NorthWall []MahjongTile
}

func NewWall(kinds []MahjongKind) Wall {
	return Wall{}
}
