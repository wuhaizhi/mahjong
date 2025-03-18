package biz

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Log("SSSSSSSSSSSSSSSSSSSSSS")
	// fmt.Println("sssssssdfadfasdfadf")
	shuffle := NewShuffle([]MahjongKind{Suit, Honor})
	fmt.Println(shuffle)
	cards, err := shuffle.Shuffle()
	if err != nil {
		t.Error(err)
	}
	//
	fmt.Printf("len: %d\n", len(cards))
	for _, v := range cards {
		fmt.Printf("card info : %s", v.Name)
	}
}
