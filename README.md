Title: Technical Challenge - Mercado Libre, Operación Fuego de Quasar
Author: Gonzalo Matías Funes
Date: 20 de Abril de 2022
Email: gonzalofunes123@gmail.com

# Operación Fuego de Quasar
|Technology | Versión |
| -- | -- |
|GoLang | 1.18 |

Para ejecutar localmente clonar el proyecto, tener instalado como mínimo la Golang 1.18. Abrir la carpeta y ejecutar desde la terminal  `go run ./main.go`. Por defecto escucha el puerto 8080

La publicación del la API se realizó en [heroku](https://www.heroku.com/) 

## EndPoints
### Local
- [POST] localhost:8080/topsecret
- [POST] localhost:8080/topsecret_split/{satellite_name}
- [GET] localhost:8080/topsecret_split
- [GET] localhost:8080 - Dummy

### Publicacion
- [POST] https://fuego-de-quasar-meli-gonzalo-f.herokuapp.com//topsecret
- [POST] https://fuego-de-quasar-meli-gonzalo-f.herokuapp.com/topsecret_split/{satellite_name}
- [GET] https://fuego-de-quasar-meli-gonzalo-f.herokuapp.com//topsecret_split
- [GET] https://fuego-de-quasar-meli-gonzalo-f.herokuapp.com/ - Dummy

## Dependencias
- [github.com/bitly/go-simplejson v0.5.0](github.com/bitly/go-simplejson)
- [github.com/gorilla/mux v1.8.0](github.com/gorilla/mux)
- [github.com/patrickmn/go-cache v2.1.0](github.com/patrickmn/go-cache)

## [Documentación](https://documenter.getpostman.com/view/6267282/UyrADwJx)


