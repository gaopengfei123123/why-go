package feature

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	deferDemo()
}

func TestDemo(他 *testing.T) {
	mp := map[string]int{"a": 1, "b": 2}
	for index, value := range mp {
		fmt.Printf("输出：index: %s, value: %d \n", index, value)
	}
}

func TestMulti(t *testing.T) {
	res, err := multipartReturn(1, 4)
	t.Logf("res: %d  err: %v", res, err)
}

func TestLoop(t *testing.T) {
	loopDemo()
}
