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
					for k := 0; k < len(SuitCards); k++ {
						cards = append(cards, MahjongTile{Type: SuitCards[k], Number: i, Name: strconv.Itoa(i) + SuitCardNames[k]})
					}
				}
			}
		case Honor:

			for j := 0; j < 4; j++ {
				for k := 0; k < len(HonorCards); k++ {
					cards = append(cards, MahjongTile{Type: HonorCards[k], Number: 0, Name: HonorCardNames[k]})
				}
			}

		case Flower:
			for j := 0; j < 4; j++ {
				for k := 0; k < len(FlowerCards); k++ {
					cards = append(cards, MahjongTile{Type: FlowerCards[k], Number: 0, Name: FlowerCardNames[k]})
				}
			}

		}
	}

	return cards, nil
}
