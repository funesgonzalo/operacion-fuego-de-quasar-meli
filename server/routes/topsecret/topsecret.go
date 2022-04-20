package topsecret

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"operacionFuegoQuasar/models"
	"operacionFuegoQuasar/services"
	"operacionFuegoQuasar/utils"
	"strings"
)

// type Satelite struct {
// 	Name     string   `json:"name"`
// 	Distance float32  `json:"distance"`
// 	Message  []string `json:"message"`
// }

// type Satelites struct {
// 	Satelites []*Satelite `json:"satellites"`
// }

// func header(res http.ResponseWriter) {
// 	res.Header().Set("Access-Control-Allow-Origin", "*")
// 	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	(res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 	res.Header().Set("Content-Type", "application/json")
// }

func TopSecretPost(res http.ResponseWriter, req *http.Request) {
	// header(res)
	utils.HeaderConfig(res)
	var bodySatelites models.Satelites

	if req.Method != "POST" {
		log.Printf("HTTP Method no soportado")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	body, err4 := ioutil.ReadAll(req.Body)
	if err4 != nil {
		log.Printf("Error al leer el body")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.Unmarshal(body, &bodySatelites)
	if err != nil {
		fmt.Println(err.Error())
	}

	var x, y, message, isValid = processRequest(bodySatelites)

	if isValid {
		result := []byte(fmt.Sprintf("{\"position\": {\"x\": %.2f,\"y\": %.2f},\"message\": \"%s\"}", x, y, message))
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {
		result := []byte(`RESPONSE CODE: 404`)
		res.WriteHeader(http.StatusNotFound)
		res.Write(result)
	}
}

func processRequest(bodySatelites models.Satelites) (float32, float32, string, bool) {
	var (
		message           string  = ""
		isValid           bool    = true
		distanceKenobi    float64 = 0
		distanceSkywalker float64 = 0
		distanceSato      float64 = 0
	)

	messages := make(map[string][]string)

	for _, s := range bodySatelites.Satelites {
		switch strings.ToLower(s.Name) {
		case "kenobi":
			distanceKenobi = float64(s.Distance)
			messages["kenobi"] = s.Message
		case "skywalker":
			distanceSkywalker = float64(s.Distance)
			messages["skywalker"] = s.Message
		case "sato":
			distanceSato = float64(s.Distance)
			messages["sato"] = s.Message
		}
	}

	kenobi := services.Point{
		X: float64(models.KenobiX),
		Y: float64(models.KenobiY),
	}

	skywalker := services.Point{
		X: float64(models.SkywalkerX),
		Y: float64(models.SkywalkerY),
	}

	sato := services.Point{
		X: float64(models.SatoX),
		Y: float64(models.SatoY),
	}

	satelites := services.NewSatellites(kenobi, skywalker, sato)
	result := satelites.GetLocation(distanceKenobi, distanceSkywalker, distanceSato)

	message = services.GetMessage(messages["kenobi"], messages["skywalker"], messages["sato"])

	if message == "" || result.X == -1 || result.Y == -1 {
		isValid = false
	}

	return float32(result.X), float32(result.Y), message, isValid
}
