package timemap

import hammingweight "github.com/lipcsei/scoring/hammingWeight"

type TimeMap struct {
	Raw [4]int `json:"raw" bson:"raw"`
}

func New(starter bool) TimeMap {
	if starter {
		return TimeMap{Raw: [4]int{-1, -1, -1, -1}}
	}
	return TimeMap{Raw: [4]int{0, 0, 0, 0}}
}

func (t TimeMap) Count() int {
	var sum, lastWordError int
	sum = 0
	for _, i := range t.Raw {
		sum += hammingweight.Of(i)
	}
	lastWordError = hammingweight.Of((t.Raw[len(t.Raw)-1] & 0xFF000000))
	return sum - lastWordError
}

func (t TimeMap) Not() TimeMap {
	for i, v := range t.Raw {
		t.Raw[i] = bitwiseNot(v)
	}
	return t
}

func bitwiseNot(b int) int {
	return 0xFFFFFFFF ^ b
}

func (t TimeMap) And(other TimeMap) TimeMap {
	for i := 0; i < len(t.Raw); i++ {
		t.Raw[i] &= other.Raw[i]
	}
	return t
}

func (t TimeMap) Or(other TimeMap) TimeMap {
	for i := 0; i < len(t.Raw); i++ {
		t.Raw[i] |= other.Raw[i]
	}
	return t
}

func (t TimeMap) Get(minute int) bool {
	if minute <= 0 || minute > 120 {
		return false
	}
	return queryBit(t.Raw[getFieldIndex(minute)], minute)
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
	for i := len(t.Raw) - 1; i > fieldIndex; i-- {
		t.Raw[i] = -1
	}
	t.Raw[fieldIndex] = t.Raw[fieldIndex] | bitwiseNot(getBitmask(minute))
	return t
}

func (t TimeMap) UnsetFrom(minute int) TimeMap {
	fieldIndex := getFieldIndex(minute)
	for i := len(t.Raw) - 1; i > fieldIndex; i-- {
		t.Raw[i] = 0
	}
	t.Raw[fieldIndex] = t.Raw[fieldIndex] & getBitmask(minute)
	return t
}

func getFieldIndex(minute int) int {
	return (minute - 1) / 32
}

func Between(from int, to int) TimeMap {
	return TimeMap{Raw: [4]int{0, 0, 0, 0}}.SetFrom(from).UnsetFrom(to + 1)
}

func Until(minute int) TimeMap {
	return TimeMap{Raw: [4]int{-1, -1, -1, -1}}.UnsetFrom(minute + 1)
}
