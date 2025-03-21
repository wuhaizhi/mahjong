package rules

import (
	"github.com/wuhaizhi/mahjong/internal/biz"
	"github.com/wuhaizhi/mahjong/internal/biz/player"
)

// 定义规则接口
type IRule interface {
	// 规则类型
	GetKind() biz.RuleType
	// 规则代码
	GetCode() string
	// 规则名称
	GetName() string
	// 规则介绍
	GetDescribe() string
	// 番数
	GetPoints() int
	// 是否符合规则
	MatchRule([]player.Player) (bool, error)
	// 优先级 0 - 100000数字越小优先级越高
	// 最初设计步长20以便后续加新的规则
	GetPriority() int
	// 不计番种 返回规则代码列表
	GetExcludePointKind() []string
}

type BaseRule struct{}

func (b *BaseRule) GetKind() biz.RuleType {
	return biz.RuleWin
}

func GetCode() string {
	return "base_rule"
}

func (b *BaseRule) GetName() string {
	return "基本规则"
}

func (b *BaseRule) GetDescribe() string {
	return `
　　1．和牌的基本牌型
　　(1) 11、123、123、123、123
　　(2) 11、123、123、123、111(1111，下同)
　　(3) 11、123、123、111、111
　　(4) 11、123、111、111、111
　　(5) 11、111、111、111、111
  `
}

func (b *BaseRule) GetPoints() int {
	return 1
}

func (b *BaseRule) MatchRule(tiles []biz.MahjongTile) (bool, error) {
	return false, nil
}

func (b *BaseRule) GetPriority() int {
	return 100000
}

func (b *BaseRule) GetExcludePointKind() []string {
	return []string{}
}
