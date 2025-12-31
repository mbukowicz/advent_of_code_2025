package day2

type isInvalidIDFunc func(int64) bool

func sumInvalidIDs(from int64, to int64, isInvalidID isInvalidIDFunc) int64 {
	sum := int64(0)
	for i := from; i <= to; i++ {
		if isInvalidID(i) {
			sum += i
		}
	}
	return sum
}
