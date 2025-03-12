package biz

import "testing"

func TestDie(t *testing.T) {
	die := NewDie()

	if die.First+die.Second > 12 {
		t.Error("骰子错误")
	}

	t.Log("die:", die.First, die.Second)
}
