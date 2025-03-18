package biz

// 牌墙
type Wall struct {
	EastWall  []MahjongCard
	SouthWall []MahjongCard
	WestWall  []MahjongCard
	NorthWall []MahjongCard
}

func NewWall(kinds []MahjongKind) Wall {
	return Wall{}
}
