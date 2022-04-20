package services

import "operacionFuegoQuasar/models"

func GetLocationService(d1, d2, d3 float32) Point {
	kenobi := Point{
		X: float64(models.KenobiX),
		Y: float64(models.KenobiY),
	}

	skywalker := Point{
		X: float64(models.SkywalkerX),
		Y: float64(models.SkywalkerY),
	}

	sato := Point{
		X: float64(models.SatoX),
		Y: float64(models.SatoY),
	}

	satelites := NewSatellites(kenobi, skywalker, sato)
	result := satelites.GetLocation(float64(d1), float64(d2), float64(d3))

	return result
}
