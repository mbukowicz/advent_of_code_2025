package day11

func countConnections(from, to string, connections map[string][]string, cache map[string]int) int {
	if from == to {
		return 1
	}

	cacheKey := from + ":" + to
	valueFromCache, existsInCache := cache[cacheKey]
	if existsInCache {
		return valueFromCache
	}

	sum := 0
	for _, device := range connections[from] {
		sum += countConnections(device, to, connections, cache)
	}

	cache[cacheKey] = sum

	return sum
}
