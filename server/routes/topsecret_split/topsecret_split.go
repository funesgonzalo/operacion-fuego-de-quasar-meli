package topsecretsplit

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
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

var CacheSplit = cache.New(5*time.Minute, 5*time.Minute)

func SetCache(key string, emp interface{}) {
	CacheSplit.Set(key, emp, cache.NoExpiration)
}

func GetCache(key string) (models.Satelite, bool) {
	var emp models.Satelite
	var found bool
	data, found := CacheSplit.Get(key)
	if found {
		emp = data.(models.Satelite)
	}
	return emp, found
}

func TopSecretSplit_Post(res http.ResponseWriter, req *http.Request) {
	utils.HeaderConfig(res)
	var bodySatelite models.Satelite
	var sateliteName string

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

	err := json.Unmarshal(body, &bodySatelite)
	if err != nil {
		fmt.Println(err.Error())
	}

	sateliteName = mux.Vars(req)["satellite_name"]
	println(sateliteName)
	bodySatelite.Name = sateliteName

	var isValid = processRequestPost_Split(bodySatelite)

	if isValid {
		result := []byte(`Se recibio con éxito la informacion del satelite ` + bodySatelite.Name)
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {
		result := []byte(`RESPONSE CODE: 404`)
		res.WriteHeader(http.StatusNotFound)
		res.Write(result)
	}
}

func processRequestPost_Split(bodySatelite models.Satelite) bool {
	var isValid bool = true

	switch strings.ToLower(bodySatelite.Name) {
	case "kenobi":
		SetCache("kenobi", bodySatelite)
	case "skywalker":
		SetCache("skywalker", bodySatelite)
	case "sato":
		SetCache("sato", bodySatelite)
	default:
		isValid = false
	}

	return isValid
}

func TopSecretSplit_Get(res http.ResponseWriter, req *http.Request) {
	utils.HeaderConfig(res)

	if req.Method != "GET" {
		log.Printf("HTTP Method no soportado")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	kenobiCache, foundKenobi := GetCache("kenobi")
	skywalkerCache, foundSkywalker := GetCache("skywalker")
	satoCache, foundSato := GetCache("sato")

	if foundKenobi && foundSkywalker && foundSato {
		// kenobi := services.Point{
		// 	X: float64(models.KenobiX),
		// 	Y: float64(models.KenobiY),
		// }

		// skywalker := services.Point{
		// 	X: float64(models.SkywalkerX),
		// 	Y: float64(models.SkywalkerY),
		// }

		// sato := services.Point{
		// 	X: float64(models.SatoX),
		// 	Y: float64(models.SatoY),
		// }

		// satelites := services.NewSatellites(kenobi, skywalker, sato)
		// result := satelites.GetLocation(float64(kenobiCache.Distance), float64(skywalkerCache.Distance), float64(satoCache.Distance))

		result := services.GetLocationService(kenobiCache.Distance, skywalkerCache.Distance, satoCache.Distance)

		message := services.GetMessage(kenobiCache.Message, skywalkerCache.Message, satoCache.Message)

		if message == "" || result.X == -1 || result.Y == -1 {
			result := []byte(`No se pudo determinar las coordenadas`)
			res.WriteHeader(http.StatusNotFound)
			res.Write(result)
		} else {
			result := []byte(fmt.Sprintf("{\"position\": {\"x\": %.2f,\"y\": %.2f},\"message\": \"%s\"}", result.X, result.Y, message))
			res.Write(result)
			res.WriteHeader(http.StatusOK)
		}
	} else {
		result := []byte(`No hay suficiente informacion para procesar la solicitud`)
		res.WriteHeader(http.StatusNotFound)
		res.Write(result)
	}

}
