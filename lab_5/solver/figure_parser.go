package solver

import "strings"

func ParseFigurePoints(figure string) (points []string) {
	figure = strings.TrimPrefix(figure, "figure=")
	return strings.Split(figure, "")
}
