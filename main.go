package main

import (
	"fmt"
	"log"
	"strconv"

	database "github.com/arturoZepeda/manejoDeCuentas/db"
	"github.com/arturoZepeda/manejoDeCuentas/extas"
	"github.com/arturoZepeda/manejoDeCuentas/gasto"
	leercsv "github.com/arturoZepeda/manejoDeCuentas/leerCSV"
)

func main() {
	/*===================Proceso BBDD=================*/
	//Inicializamos base de datos.
	gastoDB, err := database.NewGastos()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	defer gastoDB.DB.Close()

	gastos, err := gastoDB.GetGastos()
	if err != nil {
		log.Fatalf("Error al obtener los gastos: %v", err)
	}

	fmt.Println("Gastos en la BBDD: ")
	for _, g := range gastos {
		fmt.Printf("ID: %v, Descripción: %v, Monto: %v, Fecha: %v, Categoría: %v\n", g["id"], g["descripcion"], g["monto"], g["fecha"], g["categoria"])
	}
	/*==================Proceso CSV==================*/
	// genera la lectura del CSV activity y lo almacena en una matriz
	lineasTemp, err := leercsv.LeeCsv("activity.csv")
	if err != nil {
		fmt.Println(err)
	}

	// validamos  si es un amex al verificar la cabecera de los documentos.
	esAmex := leercsv.EsAmex(lineasTemp[0])
	fmt.Println(esAmex)

	// recorremos la matriz para leer cada una de las entradas
	for i, linea := range lineasTemp {

		fmt.Println("========================================")
		fmt.Println("[", i, "]")
		// Asignación de los valores para generar en el nuevo struct.
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
		// generamos nuevo struct
		gastoTemp, err := gasto.New(fecha.Format("2006-01-02"), titular, descripcion, importe)
		if err != nil {
			fmt.Println("Error al crear el gasto:", err)
			continue
		}
		gastoTemp.Output()
	}
}
