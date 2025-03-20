package biz

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	shuffle := NewShuffle([]MahjongKind{Suit, Honor})
	fmt.Println(shuffle)
	for i := 0; i < 100; i++ {
		cards, err := shuffle.Shuffle()
		if err != nil {
			t.Error(err)
		}
		//
		fmt.Printf("len: %d\n", len(cards))
		for _, v := range cards {
			fmt.Printf("%s, ", v.Name)
		}
		fmt.Println("")

	}
}
