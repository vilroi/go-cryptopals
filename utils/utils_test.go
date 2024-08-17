package utils

import (
	"testing"

	"github.com/vilroi/goutils/assert"
)

func TestCalcHamming(t *testing.T) {
	s1 := "this is a test"
	s2 := "wokka wokka!!!"

	dst := CalcHamming([]byte(s1), []byte(s2))
	assert.AssertEq(37, dst)
}
