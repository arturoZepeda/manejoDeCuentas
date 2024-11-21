package database

import (
	"database/sql"
	"log"
)

const (
	file   string = "Gasto.db"
	create string = `CREATE TABLE Gasto(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  descripcion TEXT NOT NULL,
  monto REAL NOT NULL,
  fecha TEXT NOT NULL,
  categoria TEXT NOT NULL);`
)

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

func NuevosGastos() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(create)
	if err != nil {
		return nil, err
	}
	log.Println("Base de datos creada exitosamente.")
	return db, nil
}

/*func GetGastos (*sql.DB)(error){
  res, err :=
}*/
