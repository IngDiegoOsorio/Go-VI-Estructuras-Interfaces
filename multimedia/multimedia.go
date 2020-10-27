package multimedia

type imagen struct{
	Titulo string
	Formato string
	Canales string
}

func (img *imagen) Mostrar() string {
	atributos := "Titulo: " + img.Titulo + "Formato: " + img.Formato + "Canales: " + img.Canales
	return atributos
}

type audio struct{
	Titulo string
	Formato string
	Duracion string
}

func (aud *audio) Mostrar() string {
	atributos := "Titulo: " + aud.Titulo + "Formato: " + aud.Formato + "Canales: " + aud.Duracion
	return atributos
}

type video struct{
	Titulo string
	Formato string
	Frames string
}	

func (vid *video) Mostrar() string {
	atributos := "Titulo: " + vid.Titulo + "Formato: " + vid.Formato + "Canales: " + vid.Frames
	return atributos
}