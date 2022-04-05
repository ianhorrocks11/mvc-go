package main

import (
	"mvc-go/app" // no puede haber un import /model porque estaria mal
)

func main() {
	app.StartRoute()
}





// siempre una app en go tiene que tener esto
// Hay distintos paquetes (app)
// Controllers
// Paquete de clases que viajan entre el controlador y el servicio
//service no lo vamos a usar