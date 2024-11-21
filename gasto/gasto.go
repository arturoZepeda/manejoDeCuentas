package gasto

import (
	"errors"
	"fmt"
	"time"
)

// Genero struct de Gasto para almacenarlo con posibilidad de generar json
type Gasto struct {
	FechaCreacion time.Time `json:"fecha_creacion"`
	FechaDeCompra time.Time `json:"fecha"`
	Titular       string    `json:"titular"`
	Descripcion   string    `json:"descripcion"`
	Importe       float64   `json:"importe"`
}

// Constructor para generar un nuevo struct retorna puntero
func New(fecha string, titular string, descripcion string, importe float64) (*Gasto, error) {
	if fecha == "" || descripcion == "" || importe == 0 {
		return &Gasto{}, errors.New("la fecha, descripción o importe no pueden venir vacios")
	}

	fechaDeCompra, err := time.Parse("2006-01-02", fecha)
	if err != nil {
		return &Gasto{}, errors.New("la fecha no tiene el formato correcto YYYY-MM-DD")
	}
	return &Gasto{
		time.Now(),
		fechaDeCompra,
		titular,
		descripcion,
		importe,
	}, nil
}

// metodo Output
func (gasto Gasto) Output() {
	fmt.Println("FechaCreacion:", gasto.FechaCreacion)
	fmt.Println("FechaDeCompra:", gasto.FechaDeCompra)
	fmt.Println("Titular: ", gasto.Titular)
	fmt.Println("Descripcion: ", gasto.Descripcion)
	fmt.Println("Importe: ", gasto.Importe)
}

// metodo limpiar struct:
func (gasto *Gasto) Clear() {
	gasto.FechaCreacion = time.Time{}
	gasto.FechaDeCompra = time.Time{}
	gasto.Titular = ""
	gasto.Descripcion = ""
	gasto.Importe = 0.0
}
