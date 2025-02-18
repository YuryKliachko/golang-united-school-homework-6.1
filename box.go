package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

func (b *box) idexExists(i int) bool {
	return i > len(b.shapes)-1
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("box is out of its capacity")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	err := b.idexExists(i)
	if err {
		return nil, errors.New("cannot get shape by index: index does not exists")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	err := b.idexExists(i)
	if err {
		return nil, errors.New("cannot extract shape by index: index does not exists")
	}
	var item = b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return item, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	err := b.idexExists(i)
	if err {
		return nil, errors.New("cannot replace shape by index: index does not exists")
	}
	var item = b.shapes[i]
	b.shapes[i] = shape
	return item, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, item := range b.shapes {
		sum += item.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, item := range b.shapes {
		sum += item.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var circleExist bool
	newShapes := make([]Shape, 0)
	for _, item := range b.shapes {
		_, ok := item.(*Circle)
		if ok {
			circleExist = true
		} else {
			newShapes = append(newShapes, item)
		}
	}
	b.shapes = newShapes
	if !circleExist {
		return errors.New("no Circle found in the box")
	}
	return nil
}
