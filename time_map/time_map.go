package timemap

import hammingweight "github.com/lipcsei/scoring/hammingWeight"

type TimeMap [4]int

func New(starter bool) TimeMap {
	if starter {
		return TimeMap{-1, -1, -1, -1}
	}
	return TimeMap{0, 0, 0, 0}
}

func (t TimeMap) Count() int {
	var sum, lastWordError int
	sum = 0
	for _, i := range t {
		sum += hammingweight.Of(i)
	}
	lastWordError = hammingweight.Of((t[len(t)-1] & 0xFF000000))
	return sum - lastWordError
}

func (t TimeMap) Not() TimeMap {
	for i, v := range t {
		t[i] = bitwiseNot(v)
	}
	return t
}

func bitwiseNot(b int) int {
	return 0xFFFFFFFF ^ b
}

func (t TimeMap) And(other TimeMap) TimeMap {
	for i := 0; i < len(t); i++ {
		t[i] &= other[i]
	}
	return t
}

func (t TimeMap) Or(other TimeMap) TimeMap {
	for i := 0; i < len(t); i++ {
		t[i] |= other[i]
	}
	return t
}

func (t TimeMap) Get(minute int) bool {
	if minute <= 0 || minute > 120 {
		return false
	}
	return queryBit(t[getFieldIndex(minute)], minute)
}

func queryBit(value int, minute int) bool {
	mask := 1 << getOffset(minute)
	return (value & int(mask)) != 0
}

func getBitmask(minute int) int {
	return int((1 << getOffset(minute)) - 1)
}

func getOffset(minute int) uint {
	return uint((minute - 1) % 32)
}

func (t TimeMap) SetFrom(minute int) TimeMap {

	fieldIndex := getFieldIndex(minute)
	for i := len(t) - 1; i > fieldIndex; i-- {
		t[i] = -1
	}
	t[fieldIndex] = t[fieldIndex] | bitwiseNot(getBitmask(minute))
	return t
}

func (t TimeMap) UnsetFrom(minute int) TimeMap {

	fieldIndex := getFieldIndex(minute)
	for i := len(t) - 1; i > fieldIndex; i-- {
		t[i] = 0
	}
	t[fieldIndex] = t[fieldIndex] & getBitmask(minute)
	return t
}

func getFieldIndex(minute int) int {
	return (minute - 1) / 32
}

func Between(from int, to int) TimeMap {
	return TimeMap{0, 0, 0, 0}.SetFrom(from).UnsetFrom(to + 1)
}

func Until(minute int) TimeMap {
	return TimeMap{-1, -1, -1, -1}.UnsetFrom(minute + 1)
}
