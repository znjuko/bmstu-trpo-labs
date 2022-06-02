package solver

const (
	ClassOther = iota
	ClassConvex
	ClassConcave

	TriangleCount = 3
)

type Figure struct {
	Points []Point
}

func NewFigure(pointOrder []string, points []Point) *Figure {
	pointMap := make(map[string]Point)
	for _, p := range points {
		pointMap[p.Name] = p
	}

	figurePoints := []Point{}
	for _, pointName := range pointOrder {
		point, _ := pointMap[pointName]
		figurePoints = append(figurePoints, point)
	}

	return &Figure{Points: figurePoints}
}

func (f *Figure) DetectClassType() (class int) {
	if len(f.Points) == TriangleCount {
		return f.detectTriangle()
	}

	return f.detectRectangle()
}

func (f *Figure) detectTriangle() (class int) {
	p := CountLength(f.Points[0], f.Points[1]) +
		CountLength(f.Points[1], f.Points[2]) +
		CountLength(f.Points[2], f.Points[0])
	if p/2 == CountLength(f.Points[0], f.Points[1]) ||
		p/2 == CountLength(f.Points[1], f.Points[2]) ||
		p/2 == CountLength(f.Points[2], f.Points[0]) {
		return ClassOther
	}

	return ClassConvex
}

func (f *Figure) detectRectangle() (class int) {
	var (
		a1 = f.Points[0].Y - f.Points[2].Y
		b1 = f.Points[2].X - f.Points[0].X
		a2 = f.Points[1].Y - f.Points[3].Y
		b2 = f.Points[3].X - f.Points[1].X
	)

	if (a1 / a2) != (b1 / b2) {
		return ClassOther
	}

	points := append(append([]Point{f.Points[3]}, f.Points...), f.Points[0])
	for i := 1; i <= 4; i++ {
		abX := points[i].X - points[i-1].X
		abY := points[i].Y - points[i-1].Y

		bcX := points[i+1].X - points[i].X
		bcY := points[i+1].Y - points[i].Y

		if abX*bcY-abY*bcX > 0 {
			return ClassConcave
		}
	}

	return ClassConvex
}
