// 骰子

package biz

import "math/rand"

type Die struct {
	First  int
	Second int
}

func NewDie() Die {
	die := Die{
		First:  dieRandom(),
		Second: dieRandom(),
	}
	return die
}

func dieRandom() int {
	r := rand.Int()

	awt := r % 6

	return awt + 1
}
