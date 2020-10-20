package main

import (
	"fmt"
	"./multimedias"
)

func main () {

	multimedia := multimedias.ContenidoWeb{Contenidos: []multimedias.Multimedia{}}

	fmt.Println("Multimedia")
	//Aquí menú
	var op uint = 1
	for op != 0{
		fmt.Println("")
		fmt.Println("1.- Agregar imagen")
		fmt.Println("2.- Agregar audio")
		fmt.Println("3.- Agregar video")
		fmt.Println("4.- Mostrar contenidos")
		fmt.Println("0.- Salir")

		fmt.Scan(&op)

		switch {
		case op == 1:
			var t string
			var f string
			var c int64

			fmt.Println("Titulo:")
			fmt.Scan(&t)

			fmt.Println("Formato:")
			fmt.Scan(&f)

			fmt.Println("Canales:")
			fmt.Scan(&c)

			imagen:= multimedias.Imagen{t, f, c}
			multimedia.Contenidos = append(multimedia.Contenidos, &imagen)

		case op == 2:
			var t string
			var f string
			var d int64

			fmt.Println("Titulo:")
			fmt.Scan(&t)

			fmt.Println("Formato:")
			fmt.Scan(&f)

			fmt.Println("Duracion:")
			fmt.Scan(&d)

			audio:= multimedias.Audio{t, f, d}
			multimedia.Contenidos = append(multimedia.Contenidos, &audio)

		case op == 3:
			var t string
			var f string
			var fr int64

			fmt.Println("Titulo:")
			fmt.Scan(&t)

			fmt.Println("Formato:")
			fmt.Scan(&f)

			fmt.Println("Frames:")
			fmt.Scan(&fr)

			video:= multimedias.Video{t, f, fr}
			multimedia.Contenidos = append(multimedia.Contenidos, &video)
			
		case op ==4:
			multimedia.Mostrar()

		case op == 0:
			fmt.Println("Saliendo...")

		default:
			fmt.Println("Opción inválida")
		}

	}


}