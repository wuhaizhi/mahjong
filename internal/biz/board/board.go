package board

import (
	"github.com/wuhaizhi/mahjong/internal/biz/player"
	"github.com/wuhaizhi/mahjong/internal/biz/rules"
)

// 牌桌
type Board struct {
	Players []*player.Player

	// 所有和牌类型都为true时才可以和牌
	Rules []rules.IRule
}
