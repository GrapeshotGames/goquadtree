/*
Copyright 2013 Volker Poplawski
*/

package quadtree

import "math"

// Use NewBoundingBox() to construct a BoundingBox object
type BoundingBox struct {
	MinX, MaxX, MinY, MaxY float64
}

func NewBoundingBox(xa, xb, ya, yb float64) BoundingBox {
	return BoundingBox{MinX: math.Min(xa, xb), MaxX: math.Max(xa, xb), MinY: math.Min(ya, yb), MaxY: math.Max(ya, yb)}
}

// Make BoundingBox implement the BoundingBoxer interface
func (b BoundingBox) BoundingBox() BoundingBox {
	return b
}

func (b BoundingBox) SizeX() float64 {
	return b.MaxX - b.MinX
}

func (b BoundingBox) SizeY() float64 {
	return b.MaxY - b.MinY
}

// Returns true if o intersects this
func (b BoundingBox) Intersects(o BoundingBox) bool {
	return b.MinX < o.MaxX && b.MinY < o.MaxY &&
		b.MaxX > o.MinX && b.MaxY > o.MinY
}

// Returns true if o is within this
func (b BoundingBox) Contains(o BoundingBox) bool {
	return b.MinX <= o.MinX && b.MinY <= o.MinY &&
		b.MaxX >= o.MaxX && b.MaxY >= o.MaxY
}
