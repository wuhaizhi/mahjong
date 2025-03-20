// 洗牌

package biz

import (
	"errors"
	"math/rand"
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

// 洗牌
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

// 根据牌的种类生成牌
func (s *Shuffle) GetCards() ([]MahjongTile, error) {
	cards := []MahjongTile{}

	for _, kind := range s.Kinds {
		switch kind {
		case Suit:
			for i := 1; i <= 9; i++ {
				for j := 0; j < 4; j++ {
					for k := 0; k < len(SuitCards); k++ {
						name := ""
						if SuitCards[k] == Dot {
							name = DotTiles[i-1]
						}
						if SuitCards[k] == Bamboo {
							name = BambooTiles[i-1]
						}
						if SuitCards[k] == Character {
							name = CharacterTiles[i-1]
						}
						cards = append(cards, MahjongTile{Type: SuitCards[k], Number: i, Name: name})
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
