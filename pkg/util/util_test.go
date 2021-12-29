package util

import (
	"testing"

	"github.com/Teshima-Tatsuya/GoBoy/pkg/types"

	"github.com/stretchr/testify/assert"
)

func TestUtil_Byte2Addr(t *testing.T) {
	upper := byte(0x12)
	lower := byte(0x34)

	addr := Byte2Addr(upper, lower)

	assert.Equal(t, types.Addr(0x1234), addr)
}

func TestUtil_ExtractUpperLower(t *testing.T) {
	upper := ExtractUpper(types.Addr(0x1234))
	lower := ExtractLower(types.Addr(0x1234))

	assert.Equal(t, byte(0x12), upper)
	assert.Equal(t, byte(0x34), lower)
}
