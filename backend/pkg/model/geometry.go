package model

type Point struct {
	X, Y float64
}

type PointCheck struct {
	X, Y float64
	Flag bool
}

type Line struct {
	Point1, Point2 Point
}

type ImportancePoints struct {
	Importance          string  `json:"importance"`
	ShortImportanceName string  `json:"short_importance_name"`
	Points              []Point `json:"points"`
}
