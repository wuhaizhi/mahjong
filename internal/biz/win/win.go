package win

import "github.com/wuhaizhi/mahjong/internal/biz"

// 定义和牌接口
type IWin interface {
	MatchWin([]biz.MahjongTile) (bool, error)
}
