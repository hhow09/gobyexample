package main

import (
	"fmt"
)

type PointType int

const (
	PointT PointType = iota
	Point3DT
)

type IPoint interface {
	X() float64
	Y() float64
	SetX(float64)
	SetY(float64)
	GetX() float64
	GetY() float64
}

// embed IPoint in IPoint3D will inherit all method of IPoint
type IPoint3D interface {
	IPoint
	Z() float64
	SetZ(float64)
}

type Point struct {
	x float64
	y float64
}

type Point3D struct {
	// Point struct is embedded
	Point
	z float64
}

// factory pattern for multiple struct
func NewPoint(t PointType, args ...float64) IPoint {
	switch t {
	case Point3DT:
		return &Point3D{Point: Point{x: args[0], y: args[1]}, z: args[2]}
	default:

		return &Point{x: args[0], y: args[1]}
	}
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point3D) Z() float64 {
	return p.z
}

func (p *Point3D) SetZ(z float64) {
	p.z = z
}

func main() {

	// this implements polymorphism of slice
	points := make([]IPoint, 0)

	p1 := NewPoint(PointT, 3, 4)
	p2 := NewPoint(Point3DT, 1, 2, 3)

	points = append(points, p1, p2)

	for _, p := range points {
		fmt.Printf("(%.2f %.2f)\n", p.GetX(), p.GetY())
	}
}
