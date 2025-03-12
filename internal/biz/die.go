// 骰子

package biz

import "math/rand"

type Die struct {
	First  int
	Second int
}

func (d *Die) Random() int {
	r := rand.Int()

	awt := r % 6

	return awt + 1
}

func (d *Die) FirstRandom() int {
	d.First = d.Random()
	return d.First
}

func (d *Die) SecondRandom() int {
	d.Second = d.Random()
	return d.Second
}
