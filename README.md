Title: Technical Challenge - Mercado Libre, Operación Fuego de Quasar
Author: Gonzalo Matías Funes
Date: 20 de Abril de 2022
Email: gonzalofunes123@gmail.com

# Operación Fuego de Quasar
|Technology | Versión |
| -- | -- |
|GoLang | 1.18 |

Para ejecutar localmente clonar el proyecto, tener instalado como mínimo la Golang 1.18. Abrir la carpeta y ejecutar desde la terminal  `go run ./main.go`. Por defecto escucha el puerto 8080

## EndPoints
### Local
- [POST] localhost:8080/topsecret
- [POST] localhost:8080/topsecret_split/{satellite_name}
- [GET] localhost:8080/topsecret_split
- [GET] localhost:8080 - Dummy

## Dependencias
- [github.com/bitly/go-simplejson v0.5.0](github.com/bitly/go-simplejson)
- [github.com/gorilla/mux v1.8.0](github.com/gorilla/mux)
- [github.com/patrickmn/go-cache v2.1.0](github.com/patrickmn/go-cache)
    

