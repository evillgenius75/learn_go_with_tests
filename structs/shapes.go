package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(shape Rectangle) float64 {
	return 2 * (shape.Width + shape.Height)
}

func Area(shape Rectangle) float64 {
	return shape.Width * shape.Height
}
