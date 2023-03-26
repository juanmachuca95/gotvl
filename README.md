# Gotvl

Gotvl es una simple abstracción para la traducción de mensajes y la incorporación de internacionalización <b>i18n</b> para los idiomas español e ingles `(en, es)` con el fin de abstraer la logica de la implementación. Incluyendo el middleware necesario para la ruta y validaciones por default para `gin validator`. Es importante tener en cuenta que necesita utilizar `Accept-Language` en las rutas donde este middleware sea utilizado


## Instalación

```go
go get -u github.com/juanmachuca95/gotvl
```

Puedes ver un ejemplo practico en [example]()


## Requerimientos
Este paquete solo soporta el framework de gin, por lo que solo soporta este enrutador.

## Uso

Este paquete utiliza [Gin][https://github.com/gin-gonic/gin] para la obtención de `tvl`, esto es la traducción para el idioma especificado, las validación por defecto por idioma y el localizador de mensajes para la internacionalización en dos idiomas.

1. Es necesario tener instalada la herramienta [goi18n](https://github.com/nicksnyder/go-i18n#command-goi18n) para la generación de archivos. 

2. Generar los archivos de traducción.  Necesitarás una carpeta translations donde todos los archivos `active.*.toml` estarán alojados. Si no lo tienes, puedes utilizar el <b>Makefile</b> que este repositorio provee más abajo.

Finalmente puedes utilizar este makefile para generación de los archivos de traducción que consumira el middleware.

```makefile
# Generate translations (en, es)
# Create by definitions
.PHONY: init
init:
	mkdir translations && cd translations; touch active.en.toml active.es.toml

.PHONY: gen
gen:
	cd translations && goi18n merge active.en.toml active.es.toml 

# Use the Finish command only when all translations have been completed.
.PHONY: finish
finish:
	cd translations; echo "\n" >> active.es.toml; cat translate.es.toml >> active.es.toml;

.PHONY: reset
reset: 
	cd translations; rm -rf active.es.toml translate.es.toml; touch active.en.toml

```


Establece el middleware

```go
// Accept-Language (en or es) required
r.Use(gotvl.SetInstancesTranslate)
```

Obtener la instancia
```go
tvl, err := gotvl.GetTVLContext(ctx)
```