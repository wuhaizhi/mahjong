// 洗牌

package biz

import (
	"errors"
	"math/rand"
	"strconv"
)

type Shuffle struct {
	Kinds []MahjongKind
	Cards []MahjongTile
}

func NewShuffle(kinds []MahjongKind) *Shuffle {
	return &Shuffle{
		Kinds: kinds,
	}
}

func (s *Shuffle) Shuffle() ([]MahjongTile, error) {
	cards, err := s.GetCards()
	if err != nil {
		return nil, errors.New("获取麻将牌错误")
	}
	rand.Shuffle(len(cards), func(i int, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards, nil
}

func (s *Shuffle) GetCards() ([]MahjongTile, error) {
	cards := []MahjongTile{}

	for _, kind := range s.Kinds {
		switch kind {
		case Suit:
			for i := 1; i <= 9; i++ {
				for j := 0; j < 4; j++ {
					cards = append(cards, MahjongTile{Type: Dot, Number: i, Name: strconv.Itoa(i) + "筒"})
					cards = append(cards, MahjongTile{Type: Bamboo, Number: i, Name: strconv.Itoa(i) + "条"})
					cards = append(cards, MahjongTile{Type: Character, Number: i, Name: strconv.Itoa(i) + "万"})
				}
			}
		case Honor:
			for j := 0; j < 4; j++ {
				cards = append(cards, MahjongTile{Type: East, Number: 0, Name: "东"})
				cards = append(cards, MahjongTile{Type: South, Number: 0, Name: "南"})
				cards = append(cards, MahjongTile{Type: Wast, Number: 0, Name: "西"})
				cards = append(cards, MahjongTile{Type: North, Number: 0, Name: "北"})
				cards = append(cards, MahjongTile{Type: Red, Number: 0, Name: "中"})
				cards = append(cards, MahjongTile{Type: Green, Number: 0, Name: "发"})
				cards = append(cards, MahjongTile{Type: White, Number: 0, Name: "白"})
			}

		case Flower:
			for j := 0; j < 4; j++ {
				cards = append(cards, MahjongTile{Type: Spring, Number: 0, Name: "春"})
				cards = append(cards, MahjongTile{Type: Summer, Number: 0, Name: "夏"})
				cards = append(cards, MahjongTile{Type: Autumn, Number: 0, Name: "秋"})
				cards = append(cards, MahjongTile{Type: Winter, Number: 0, Name: "冬"})
				cards = append(cards, MahjongTile{Type: Plum, Number: 0, Name: "梅"})
				cards = append(cards, MahjongTile{Type: Orchid, Number: 0, Name: "兰"})
				cards = append(cards, MahjongTile{Type: Bamboo01, Number: 0, Name: "竹"})
				cards = append(cards, MahjongTile{Type: Chrysanthemum, Number: 0, Name: "菊"})
			}

		}
	}

	return cards, nil
}
