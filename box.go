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

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("error out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.shapesCapacity < i {
		return nil, errors.New("index went lower than i")
	}
	if len(b.shapes) <= i || i < 0 {
		return nil, errors.New("shape by index doesn't exist")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.shapesCapacity < i {
		return nil, errors.New("index went lower than i")
	}
	if len(b.shapes) <= i || i < 0 {
		return nil, errors.New("shape by index doesn't exist")
	}
	newShape := b.shapes[i]
	b.shapes = b.shapes[:i+copy(b.shapes[i:], b.shapes[i+1:])]
	return newShape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.shapesCapacity < i {
		return nil, errors.New("index went lower than i")
	}
	if len(b.shapes) <= i || i < 0 {
		return nil, errors.New("shape by index doesn't exist")
	}
	newShape := b.shapes[i]
	b.shapes[i] = shape
	return newShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var list float64
	for i := 0; i < len(b.shapes); i++ {
		list += b.shapes[i].CalcPerimeter()
	}
	return list
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var list float64
	for i := 0; i < len(b.shapes); i++ {
		list += b.shapes[i].CalcArea()
	}
	return list
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShape := make([]Shape, 0,10)
	for i := 0; i < len(b.shapes); i++ {
		if _, exist := b.shapes[i].(*Circle); !exist {
			newShape = append(newShape, b.shapes[i])
		}
	}
	if len(b.shapes)==len(newShape){
		return errors.New("exist in the list")
	}
	b.shapes=newShape
	return nil
}
