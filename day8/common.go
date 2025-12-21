package day8

type Point struct {
	X, Y, Z int
}

type Distance struct {
	A, B *Point
	dist int
}

func distanceSquared(a, b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	return dx*dx + dy*dy + dz*dz
}
