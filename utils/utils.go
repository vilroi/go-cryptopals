package utils

import (
	"math/bits"

	"github.com/vilroi/goutils/assert"
)

func CalcHamming(x, y []byte) int {
	assert.AssertEq(len(x), len(y))

	var dist int
	for i := 0; i < len(x); i++ {
		t := x[i] ^ y[i]
		dist += bits.OnesCount8(t)
	}

	return dist
}
