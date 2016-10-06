package timemap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMethod(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(TimeMap{}, New(true))
	assert.Equal(TimeMap{0, 0, 0, 0}, New(false))
}

func TestBitCountZero(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint32(0), New(false).Count())
}

func TestBitCountNormal(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint32(16), TimeMap{15, 15, 15, 15}.Count())
}

func TestBitwiseNotFunction(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint32(0xFFF0FFF0), bitwiseNot(0xF000F))
}

func TestNotMethod(t *testing.T) {
	assert := assert.New(t)
	expected := TimeMap{0xFFFFFFFD, 0xFFFFFFFE, 0xFFFFFFFF, 0xFFFFFFFD}
	actual := TimeMap{0x2, 0x1, 0x0, 0x2}.Not()
	assert.Equal(expected, actual)
}

func TestAndMethod(t *testing.T) {
	assert := assert.New(t)
	a := TimeMap{0x2, 0x1}
	b := TimeMap{0x2, 0x5, 0x3, 0x5}
	expected := TimeMap{0x2, 0x1}
	assert.Equal(expected, a.And(b))
}

func TestOrMethod(t *testing.T) {
	assert := assert.New(t)
	a := TimeMap{0x2, 0x1}
	b := TimeMap{0x2, 0x5, 0x3, 0x5}
	expected := TimeMap{0x2, 0x5, 0x3, 0x5}

	assert.Equal(expected, a.Or(b))
}

func TestGetMethod(t *testing.T) {
	assert := assert.New(t)
	assert.False(TimeMap{}.Get(0))
	assert.False(TimeMap{}.Get(121))

	fmt.Println(49 % 32)

}

func TestGetBitmaskFunction(t *testing.T) {
	assert := assert.New(t)
	expected := uint32(0x1FFFF)
	assert.Equal(expected, getBitmask(82))
}

func TestGetFieldIndexFunction(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint32(0), getFieldIndex(12))
	assert.Equal(uint32(1), getFieldIndex(40))
	assert.Equal(uint32(2), getFieldIndex(70))
	assert.Equal(uint32(3), getFieldIndex(100))
}

// func TestQueryBitFunction(t *testing.T) {
// 	assert := assert.New(t)
// 	// assert.True(queryBit(0, 42))
// }

// func TestSet
