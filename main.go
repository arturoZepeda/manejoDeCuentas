package main

import (
	"fmt"
	"strconv"

	"github.com/arturoZepeda/manejoDeCuentas/extas"
	"github.com/arturoZepeda/manejoDeCuentas/gasto"
	leercsv "github.com/arturoZepeda/manejoDeCuentas/leerCSV"
)

func main() {
	lineasTemp, err := leercsv.LeeCsv("activity.csv")
	if err != nil {
		fmt.Println(err)
	}
	esAmex := leercsv.EsAmex(lineasTemp[0])
	fmt.Println(esAmex)
	for i, linea := range lineasTemp {
		fmt.Println("========================================")
		fmt.Println("[", i, "]")
		fecha, err := extas.ParseFechaEspanol(linea[0])
		if err != nil {
			fmt.Println("error al convertir la fecha")
		}
		titular := linea[3]
		descripcion := linea[2]
		importeStr := linea[5]
		importe, err := strconv.ParseFloat(importeStr, 64)
		if err != nil {
			fmt.Println("Problemas al parsear el importe.", err)
		}
		fmt.Println(fecha, titular, descripcion, importe)
		gastoTemp, err := gasto.New(fecha.Format("2006-01-02"), titular, descripcion, importe)
		if err != nil {
			fmt.Println("Error al crear el gasto:", err)
			continue
		}
		fmt.Println(gastoTemp)
	}
}
