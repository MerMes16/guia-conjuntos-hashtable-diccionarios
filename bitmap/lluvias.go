package bitmap

import "untref-ayp2/guia-conjuntos-hashes-diccionarios/set"

// Definimos el tipo Mes que puede tomar valores enteros entre 0 y 11
// podemos usar directamente el nombre del Mes en vez de usar el
// valor numérico, por ejemplo: enero, febrero, marzo, etc.
type Mes uint8

const (
	enero Mes = iota
	febrero
	marzo
	abril
	mayo
	junio
	julio
	agosto
	septiembre
	octubre
	noviembre
	diciembre
)

// valida si el Mes y el dia son validos
func validaFecha(d uint8, m Mes) bool {
	switch m {
	case abril, junio, septiembre, noviembre:
		return d >= 1 && d <= 30
	case febrero:
		return d >= 1 && d <= 29
	default: // enero, marzo, mayo, julio, agosto, octubre, diciembre
		return d >= 1 && d <= 31
	}
}

// Lluvias es una estructura que representa la cantidad de lluvias caídas en un año
type Lluvias struct {
	meses [12]*BitMap
}

// NewLluvias crea un nuevo registro de lluvias
func NewLluvias() *Lluvias {
	var meses [12]*BitMap
	for i := range 12 {
		meses[i] = NewBitMap()
	}
	return &Lluvias{meses}
}

// Registrar lluvia registra un dia del año en que llovió
// Si la fecha es inválida, no hace nada
func (l *Lluvias) Registrar(d uint8, m Mes) {
	if validaFecha(d, m) {
		l.meses[m].On(d)
	}
}

// Desregistrar lluvia desregistra un dia del año en que llovió
func (l *Lluvias) Desregistrar(d uint8, m Mes) {
	if validaFecha(d, m) {
		l.meses[m].Off(d)
	}
}

// Llovió devuelve si llovió en un día del año
// Si la fecha es inválida, devuelve false
func (l *Lluvias) Llovio(d uint8, m Mes) bool {
	var llovio bool
	if validaFecha(d, m) {
		llovio, _ = l.meses[m].IsOn(d)
	}
	return llovio
}

// ListarMes lista los días en que llovió en un Mes
func (l *Lluvias) ListarMes(m Mes) *set.ListSet[uint8] {
	set := set.NewListSet[uint8]()
	for i := range 32 {
		isOn, _ := l.meses[m].IsOn(uint8(i + 1))
		if isOn {
			set.Add(uint8(i + 1))
		}
	}
	return set
}

// Dado dos meses, listar los días que llovieron en ambos
func (l *Lluvias) ListarDiasEnAmbosMeses(m1 Mes, m2 Mes) *set.ListSet[uint8] {
	set1 := l.ListarMes(m1)
	set2 := l.ListarMes(m2)
	return set1.Intersection(set2)
}
