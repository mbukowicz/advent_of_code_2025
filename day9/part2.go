package day9

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPointOnLine(point Point, line Line) bool {
	// point on horizontal line
	if point.X == line.Start.X && line.Start.X == line.End.X && min(line.Start.Y, line.End.Y) <= point.Y && point.Y <= max(line.Start.Y, line.End.Y) {
		return true
	}

	// point on vertical line
	if point.Y == line.Start.Y && line.Start.Y == line.End.Y && min(line.Start.X, line.End.X) <= point.X && point.X <= max(line.Start.X, line.End.X) {
		return true
	}

	return false
}

func isPointCrossingRayCastingLineOnTheLeft(point Point, line Line) bool {
	lineStartsAbovePoint := line.Start.Y > point.Y
	lineEndsAbovePoint := line.End.Y > point.Y
	if lineStartsAbovePoint == lineEndsAbovePoint {
		return false
	}

	if line.Start.X != line.End.X {
		// this must be a horizontal line because line.Start.Y != line.End.Y (see above condition)
		panic("Not a horizontal line")
	}

	// in this simplified ray casting algorithm we only consider horizontal lines so we
	// can simply check if the imagined ray starting in point crosses the line from the right side
	lineOnTheLeft := line.End.X < point.X
	return lineOnTheLeft
}

func isPointInsidePolygon(point Point, polygon []Line) bool {
	inside := false
	for _, line := range polygon {
		if isPointOnLine(point, line) {
			return true
		}
		if isPointCrossingRayCastingLineOnTheLeft(point, line) {
			inside = !inside
		}
	}
	return inside
}

func allRectangleCornersInsidePolygon(corners [4]Point, polygon []Line) bool {
	for _, point := range corners {
		if !isPointInsidePolygon(point, polygon) {
			return false
		}
	}
	return true
}

func isLineCrossingRectangle(line Line, topLeftCorner Point, bottomRightCorner Point) bool {
	if line.Start.Y == line.End.Y {
		// horizontal line
		lineY := line.Start.Y
		if topLeftCorner.Y < lineY && lineY < bottomRightCorner.Y {
			return topLeftCorner.X < max(line.Start.X, line.End.X) && bottomRightCorner.X > min(line.Start.X, line.End.X)
		}
	} else if line.Start.X == line.End.X {
		// vertical line
		lineX := line.Start.X
		if topLeftCorner.X < lineX && lineX < bottomRightCorner.X {
			return topLeftCorner.Y < max(line.Start.Y, line.End.Y) && bottomRightCorner.Y > min(line.Start.Y, line.End.Y)
		}
	} else {
		panic("Neither horizontal nor vertical line")
	}
	return false
}

func rectangleIntersectsWithPolygon(topLeftCorner Point, bottomRightCorner Point, polygon []Line) bool {
	for _, line := range polygon {
		if isLineCrossingRectangle(line, topLeftCorner, bottomRightCorner) {
			return true
		}
	}
	return false
}

func isValidRectangle(point1 Point, point2 Point, polygon []Line) bool {
	minX := min(point1.X, point2.X)
	maxX := max(point1.X, point2.X)
	minY := min(point1.Y, point2.Y)
	maxY := max(point1.Y, point2.Y)
	topLeftCorner := Point{X: minX, Y: minY}
	topRightCorner := Point{X: maxX, Y: minY}
	bottomRightCorner := Point{X: maxX, Y: maxY}
	bottomLeftCorner := Point{X: minX, Y: maxY}
	corners := [4]Point{topLeftCorner, topRightCorner, bottomRightCorner, bottomLeftCorner}
	if !allRectangleCornersInsidePolygon(corners, polygon) {
		return false
	}
	if rectangleIntersectsWithPolygon(topLeftCorner, bottomRightCorner, polygon) {
		return false
	}

	return true
}

func Part2() {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	points := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			panic("invalid point format")
		}
		x, y := 0, 0
		fmt.Sscanf(parts[0], "%d", &x)
		fmt.Sscanf(parts[1], "%d", &y)
		point := Point{X: x, Y: y}
		points = append(points, point)
	}

	polygon := []Line{}
	for i := 0; i < len(points)-1; i++ {
		polygon = append(polygon, Line{Start: points[i], End: points[i+1]})
	}
	polygon = append(polygon, Line{Start: points[len(points)-1], End: points[0]})

	maxSize := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			width := abs(points[i].X-points[j].X) + 1
			height := abs(points[i].Y-points[j].Y) + 1
			size := width * height
			if size > maxSize && isValidRectangle(points[i], points[j], polygon) {
				fmt.Println("New max size found:", size, "with points", points[i], "and", points[j])
				maxSize = size
			}
		}
	}

	fmt.Printf("Result: %d\n", maxSize)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
	}
}
