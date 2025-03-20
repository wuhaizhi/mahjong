package rules

import (
	"github.com/wuhaizhi/mahjong/internal/biz"
)

// 定义宝牌接口
type IRule interface {
	// 规则名称
	GetName() string
	// 规则介绍
	GetDescribe() string
	// 分值
	GetScore() int
	// 是否符合规则
	MatchRule([]biz.MahjongTile) (bool, error)
}
