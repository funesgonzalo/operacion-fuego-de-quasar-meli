package models

const KenobiX = -500
const KenobiY = -200

const SkywalkerX = 100
const SkywalkerY = -100

const SatoX = 500
const SatoY = 100

type Satelite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Satelites struct {
	Satelites []*Satelite `json:"satellites"`
}
