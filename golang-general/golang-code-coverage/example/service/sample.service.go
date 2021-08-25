package service

import (
	"github.com/aasisodiya/go/codecoverage/constants"
	"math"
)

// Area Interface
type CubeOp interface {
	SurfaceArea() (area int)
	Volume() (volume int)
}

// Cube struct
type Cube struct {
	side int
}

// NewCube to Initialize a cube
func NewCube(size int) *Cube {
	return &Cube{side: size}
}

// SurfaceArea will Calculate and Returns Surface Area
func (cube Cube) SurfaceArea() (area int) {
	return constants.NumberOfFacesInCube * cube.side * cube.side
}

// Volume will Calculate and Returns Volume
func (cube Cube) Volume() (volume int) {
	return int(math.Pow(float64(cube.side), 3))
}
