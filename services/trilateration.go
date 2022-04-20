package services

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Satellites struct {
	Kenobi    Point
	Skywalker Point
	Sato      Point
}

func (s *Satellites) GetLocation(d1, d2, d3 float64) (coordinates Point) {
	var shipCoordinates Point

	//unit vector in a direction from point1 to point 2
	kenobiSkywalkerDistance := math.Pow(math.Pow(s.Skywalker.X-s.Kenobi.X, 2)+math.Pow(s.Skywalker.Y-s.Kenobi.Y, 2), 0.5)

	ex := Point{
		X: (s.Skywalker.X - s.Kenobi.X) / kenobiSkywalkerDistance,
		Y: (s.Skywalker.Y - s.Kenobi.Y) / kenobiSkywalkerDistance,
	}
	aux := Point{
		X: s.Sato.X - s.Kenobi.X,
		Y: s.Sato.Y - s.Kenobi.Y,
	}

	//signed magnitude of the x component
	i := ex.X*aux.X + ex.Y*aux.Y

	//the unit vector in the y direction.
	aux2 := Point{
		X: s.Sato.X - s.Kenobi.X - i*ex.X,
		Y: s.Sato.Y - s.Kenobi.Y - i*ex.Y,
	}
	ey := Point{
		X: aux2.X / normalize(aux2),
		Y: aux2.Y / normalize(aux2),
	}

	//the signed magnitude of the y component
	j := ey.X*aux.X + ey.Y*aux.Y

	//coordinates
	x := (math.Pow(d1, 2) - math.Pow(d2, 2) + math.Pow(kenobiSkywalkerDistance, 2)) / (2 * kenobiSkywalkerDistance)
	y := (math.Pow(d1, 2)-math.Pow(d3, 2)+math.Pow(i, 2)+math.Pow(j, 2))/(2*j) - (i*x)/j

	//result coordinates
	shipCoordinates.X = s.Kenobi.X + x*ex.X + y*ey.X
	shipCoordinates.Y = s.Kenobi.Y + x*ex.Y + y*ey.Y

	return shipCoordinates
}

func normalize(point Point) float64 {
	return math.Pow(math.Pow(point.X, 2)+math.Pow(point.Y, 2), .5)
}

func NewSatellites(
	Kenobi Point,
	Skywalker Point,
	Sato Point,
) *Satellites {
	return &Satellites{
		Kenobi:    Kenobi,
		Skywalker: Skywalker,
		Sato:      Sato,
	}
}
