package main

import (
	"fmt"
	"math"
)

func main() {
	type Point struct {
		x int
		y int
	}

	type Circle struct {
		center Point
		radius int
	}
	var circle1, circle2 Circle
	var point Point

	fmt.Scan(&circle1.center.x, &circle1.center.y, &circle1.radius)
	fmt.Scan(&circle2.center.x, &circle2.center.y, &circle2.radius)
	fmt.Scan(&point.x, &point.y)

	distanceToCircle1 := math.Sqrt(math.Pow(float64(point.x-circle1.center.x), 2) + math.Pow(float64(point.y-circle1.center.y), 2))
	distanceToCircle2 := math.Sqrt(math.Pow(float64(point.x-circle2.center.x), 2) + math.Pow(float64(point.y-circle2.center.y), 2))
	
	if distanceToCircle1 < float64(circle1.radius) && distanceToCircle2 < float64(circle2.radius) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if distanceToCircle1 < float64(circle1.radius) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if distanceToCircle2 < float64(circle2.radius) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}