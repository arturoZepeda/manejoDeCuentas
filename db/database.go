package database

import (
	"database/sql"
	"log"
)

const file string = "Gasto.db"

func Conectar() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Conexion exitosa a la base de datos: ", file)
	return db, nil
}
