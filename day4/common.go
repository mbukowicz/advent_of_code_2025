package day4

var directions = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func countAdjacentOccupied(rows [][]bool, r, c int) int {
	count := 0
	for _, dir := range directions {
		newR := r + dir[0]
		newC := c + dir[1]
		if newR >= 0 && newR < len(rows) && newC >= 0 && newC < len(rows[0]) {
			if rows[newR][newC] {
				count++
			}
		}
	}
	return count
}
