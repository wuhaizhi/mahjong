// 洗牌

package biz

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Shuffle struct {
	Kinds []MahjongKind
	Cards []MahjongCard
}

func NewShuffle(kinds []MahjongKind) *Shuffle {
	return &Shuffle{
		Kinds: kinds,
	}
}

func (s *Shuffle) Shuffle() ([]MahjongCard, error) {
	cards, err := s.GetCards()
	if err != nil {
		return nil, err
	}
	cardsLen := len(cards)

	shuffleCards := []MahjongCard{}

	cardsUse := map[int]struct{}{}

	for i := 0; i < cardsLen; i++ {
		index := rand.Intn(cardsLen - 1)

		fmt.Println("index1:", index)
		_, ok := cardsUse[index]
		for ok {
			if index >= cardsLen-1 {
				index = 0
			} else {
				index++
			}
			_, ok = cardsUse[index]

		}

		cardsUse[index] = struct{}{}
		shuffleCards = append(shuffleCards, cards[index])
	}

	return shuffleCards, nil
}

func (s *Shuffle) GetCards() ([]MahjongCard, error) {
	cards := []MahjongCard{}

	for _, kind := range s.Kinds {
		switch kind {
		case Suit:
			for i := 1; i <= 9; i++ {
				for j := 0; j < 4; j++ {
					cards = append(cards, MahjongCard{Type: Dot, Number: i, Name: strconv.Itoa(i) + "筒"})
					cards = append(cards, MahjongCard{Type: Bamboo, Number: i, Name: strconv.Itoa(i) + "条"})
					cards = append(cards, MahjongCard{Type: Character, Number: i, Name: strconv.Itoa(i) + "万"})
				}
			}
		case Honor:
			for j := 0; j < 4; j++ {
				cards = append(cards, MahjongCard{Type: East, Number: 0, Name: "东"})
				cards = append(cards, MahjongCard{Type: South, Number: 0, Name: "南"})
				cards = append(cards, MahjongCard{Type: Wast, Number: 0, Name: "西"})
				cards = append(cards, MahjongCard{Type: North, Number: 0, Name: "北"})
				cards = append(cards, MahjongCard{Type: Red, Number: 0, Name: "中"})
				cards = append(cards, MahjongCard{Type: Green, Number: 0, Name: "发"})
				cards = append(cards, MahjongCard{Type: White, Number: 0, Name: "白"})
			}

		case Flower:
			for j := 0; j < 4; j++ {
				cards = append(cards, MahjongCard{Type: Spring, Number: 0, Name: "春"})
				cards = append(cards, MahjongCard{Type: Summer, Number: 0, Name: "夏"})
				cards = append(cards, MahjongCard{Type: Autumn, Number: 0, Name: "秋"})
				cards = append(cards, MahjongCard{Type: Winter, Number: 0, Name: "冬"})
				cards = append(cards, MahjongCard{Type: Plum, Number: 0, Name: "梅"})
				cards = append(cards, MahjongCard{Type: Orchid, Number: 0, Name: "兰"})
				cards = append(cards, MahjongCard{Type: Bamboo01, Number: 0, Name: "竹"})
				cards = append(cards, MahjongCard{Type: Chrysanthemum, Number: 0, Name: "菊"})
			}

		}
	}

	return cards, nil
}
