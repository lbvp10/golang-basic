package main

import "fmt"

type Pago struct {
	id      int64
	monto   float64
	cliente string
	propina float64
}

func main() {
	p := Pago{
		id:      1,
		cliente: "Juan",
		monto:   2000,
	}
	p.imprimir()
	fmt.Println(p.calcularIva(0.19))

	adicionarPropina(&p)
	p.imprimir()
}
func (p *Pago) imprimir() {
	fmt.Println(p)
}

func (p *Pago) calcularIva(iva float32) float32 {
	return float32(p.monto) * iva
}

func adicionarPropina(p *Pago) {
	p.propina = float64(p.monto) * 0.03
}
