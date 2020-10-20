package multimedias

import "fmt"

// - Imagen: titulo, formato, canales
type Imagen struct {
	Titulo string
	Formato string
	Canales int64
}

// - Audio: titulo, formato, duracion (seg)
type Audio struct {
	Titulo string
	Formato string
	Duracion int64
}
// - Video: titulo, formato, frames
type Video struct {
	Titulo string
	Formato string
	Frames int64
}

func (i *Imagen) Mostrar() {
	fmt.Println("Titulo:", i.Titulo)
	fmt.Println("Formato:", i.Formato)
	fmt.Println("Canales:", i.Canales)
}

func (a *Audio) Mostrar() {
	fmt.Println("Titulo:", a.Titulo)
	fmt.Println("Formato:", a.Formato)
	fmt.Println("Duracion:", a.Duracion)
}

func (v *Video) Mostrar() {
	fmt.Println("Titulo:", v.Titulo)
	fmt.Println("Formato:", v.Formato)
	fmt.Println("Frames:", v.Frames)
}

//Interfaz
type Multimedia interface{
	Mostrar()
}

//Estructura ContendidoWeb con slice de Multimedias
type ContenidoWeb struct {
	Contenidos []Multimedia
}

func (cw *ContenidoWeb) Mostrar() {

	for _, c := range cw.Contenidos{
		c.Mostrar()
		fmt.Println("")
	}
}



