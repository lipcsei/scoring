package timemap

import hammingweight "github.com/lipcsei/scoring/hammingWeight"

type TimeMap [4]uint32

func New(starter bool) TimeMap {
	if starter {
		return TimeMap{}
	}
	return TimeMap{0, 0, 0, 0}
}

func (t TimeMap) Count() uint32 {
	var sum, lastWordError uint32
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

func bitwiseNot(b uint32) uint32 {
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

func (t TimeMap) Get(minute uint32) bool {
	if minute <= 0 || minute > 120 {
		return false
	}

	return queryBit(t[getFieldIndex(minute)], minute)
}

func queryBit(value uint32, minute uint32) bool {
	return (value & getOffset(minute)) != 0
}

func getBitmask(minute uint32) uint32 {
	return getOffset(minute-1) - 1
}

func getOffset(minute uint32) uint32 {
	return 1 << (minute - 1) % 32
}

func (t TimeMap) SetFrom(minute uint32) {
	fieldIndex := getFieldIndex(minute)
	for i := uint32(len(t) - 1); i > fieldIndex; i-- {
		t[i] = 0
	}
	t[fieldIndex] = t[fieldIndex] | bitwiseNot(getBitmask(minute))
}

func (t TimeMap) UnsetFrom(minute uint32) {
	fieldIndex := getFieldIndex(minute)
	for i := uint32(len(t) - 1); i > fieldIndex; i-- {
		t[i] = 0
	}
	t[fieldIndex] = t[fieldIndex] & getBitmask(minute)
}

func getFieldIndex(minute uint32) uint32 {
	return (minute - 1) / 32
}
