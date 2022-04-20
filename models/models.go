package models

// type Satelite struct {
// 	Longitude float32
// 	Latitude  float32
// }

const KenobiX = -500
const KenobiY = -200

const SkywalkerX = 100
const SkywalkerY = -100

const SatoX = 500
const SatoY = 100

// var (
// 	Kenobi    = Satelite{kenobiX, kenobiY}
// 	Skywalker = Satelite{skywalkerX, skywalkerY}
// 	Sato      = Satelite{satoX, satoY}
// )

// var (
// 	Kenobi    = Point{kenobiX, kenobiY}
// 	Skywalker = Point{skywalkerX, skywalkerY}
// 	Sato      = Point{satoX, satoY}
// )

// type Point struct {
// 	X float64
// 	Y float64
// }

// type Satellites struct {
// 	Kenobi    Point
// 	Skywalker Point
// 	Sato      Point
// }

type Satelite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Satelites struct {
	Satelites []*Satelite `json:"satellites"`
}
