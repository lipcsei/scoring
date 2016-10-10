package timemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMethod(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(TimeMap{-1, -1, -1, -1}, New(true))
	assert.Equal(TimeMap{0, 0, 0, 0}, New(false))
}

func TestBitwiseNotFunction(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0xFFF0FFF0, bitwiseNot(0xF000F))
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

func TestGetBitmaskFunction(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0x1FFFF, getBitmask(82))
}

func TestGetFieldIndexFunction(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, getFieldIndex(12))
	assert.Equal(1, getFieldIndex(40))
	assert.Equal(2, getFieldIndex(70))
	assert.Equal(3, getFieldIndex(100))
}

func TestGetOffsetMethod(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(uint(17), getOffset(82))
	assert.Equal(uint(28), getOffset(29))
}

func TestQueryBitFunction(t *testing.T) {
	assert := assert.New(t)
	assert.True(queryBit(-1, 29))
	assert.True(queryBit(67108863, 68))
}

func TestStarterReturnsTrueForEveryMinute(t *testing.T) {
	assert := assert.New(t)

	timemap := New(true)
	for i := 1; i < 120; i++ {
		assert.True(timemap.Get(i))
	}
}

func TestNotStarterReturnsFalseForEveryMinute(t *testing.T) {
	assert := assert.New(t)

	timemap := New(false)
	for i := 1; i < 120; i++ {
		assert.False(timemap.Get(i))
	}
}
func TestGetReturnsFalseForInvalidValues(t *testing.T) {
	assert := assert.New(t)
	timemap := New(true)
	assert.False(timemap.Get(0))
	assert.False(timemap.Get(125))
}

func TestGetReturnsCorrectValues(t *testing.T) {
	assert := assert.New(t)
	timemap := TimeMap{15, 0, 15, 0}

	for i := 1; i <= 4; i++ {
		assert.True(timemap.Get(i))
	}
	for i := 5; i <= 64; i++ {
		assert.False(timemap.Get(i))
	}
	for i := 65; i <= 68; i++ {
		assert.True(timemap.Get(i))
	}
	for i := 69; i <= 120; i++ {
		assert.False(timemap.Get(i))
	}
}

func TestSetThenGetLower(t *testing.T) {
	assert := assert.New(t)
	timemap := TimeMap{0, 0, 0, 0}
	timemap = timemap.SetFrom(10)
	for i := 1; i <= 9; i++ {
		assert.False(timemap.Get(i))
	}

	for i := 10; i <= 120; i++ {
		assert.True(timemap.Get(i))
	}
}

func TestUnsetThenGetLower(t *testing.T) {
	assert := assert.New(t)
	timemap := TimeMap{-1, -1, -1, -1}
	timemap = timemap.UnsetFrom(10)
	for i := 1; i <= 9; i++ {
		assert.True(timemap.Get(i))
	}

	for i := 10; i <= 120; i++ {
		assert.False(timemap.Get(i))
	}
}

func TestUnsetThenGetUpper(t *testing.T) {
	assert := assert.New(t)
	timemap := TimeMap{-1, -1, -1, -1}
	timemap = timemap.UnsetFrom(75)
	for i := 1; i <= 74; i++ {
		assert.True(timemap.Get(i))
	}

	for i := 75; i <= 120; i++ {
		assert.False(timemap.Get(i))
	}
}

func TestBitCountZero(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, New(false).Count())
}

func TestBitCountNormal(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(16, TimeMap{15, 15, 15, 15}.Count())
}

func TestBitCountOverflowStarter(t *testing.T) {
	assert := assert.New(t)
	timemap := New(true)
	assert.Equal(120, timemap.Count())
}

func TestBitCountOverflowAfterSet(t *testing.T) {
	assert := assert.New(t)
	timemap := TimeMap{0, 0, 0, 0}
	timemap = timemap.SetFrom(71)
	assert.Equal(50, timemap.Count())
}

func TestUntilLower(t *testing.T) {
	assert := assert.New(t)
	timemap := Until(25)
	for i := 1; i <= 25; i++ {
		assert.True(timemap.Get(i))
	}

	for i := 26; i <= 120; i++ {
		assert.False(timemap.Get(i))
	}
}

func TestUntilHigher(t *testing.T) {
	assert := assert.New(t)
	timemap := Until(75)
	for i := 1; i <= 75; i++ {
		assert.True(timemap.Get(i))
	}

	for i := 76; i <= 120; i++ {
		assert.False(timemap.Get(i))
	}
}

func TestUntilMax(t *testing.T) {
	assert := assert.New(t)
	timemap := Until(120)
	for i := 1; i <= 120; i++ {
		assert.True(timemap.Get(i))
	}
}

func TestBetweenLower(t *testing.T) {
	assert := assert.New(t)
	timemap := Between(25, 45)
	for i := 1; i < 25; i++ {
		assert.False(timemap.Get(i))
	}

	for i := 25; i <= 45; i++ {
		assert.True(timemap.Get(i))
	}

	for i := 46; i <= 120; i++ {
		assert.False(timemap.Get(i))
	}

}
