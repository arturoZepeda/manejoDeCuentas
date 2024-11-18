package leercsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func LeeCsv(fichero string) ([][]string, error) {
	file, err := os.Open(fichero)
	if err != nil {
		return nil, errors.New("error al leer el fichero: ")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	lineas, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error leyendo el fichero.")
		return nil, errors.New("error al leer las lineas")
	}
	return lineas, nil
}

func EsAmex(cabecera []string) bool {
	if cabecera[0] == "Fecha" && cabecera[1] == "Fecha de Compra" && cabecera[2] == "Descripción" && cabecera[3] == "Titular de la Tarjeta" {
		return true
	}
	return false
}

/*
Fecha
Fecha de Compra
Descripción
Titular de la Tarjeta
Cuenta
Importe
*/
