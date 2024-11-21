
# Manejo de Cuentas

Este proyecto es una herramienta en Go para la gestión de gastos personales, lectura de archivos CSV con datos financieros y manipulación de una base de datos SQLite para almacenar y recuperar información.

## Estructura del Proyecto

```
.
├── Gasto.db                # Base de datos SQLite (archivo raíz, inicializable automáticamente)
├── activity.csv            # Archivo CSV de ejemplo para pruebas
├── db/
│   ├── database.go         # Gestión de la base de datos y funciones relacionadas
│   └── gastos.db           # Base de datos alternativa en carpeta db/
├── extas/
│   └── extras.go           # Funciones auxiliares (e.g., parseo de fechas)
├── gasto/
│   └── gasto.go            # Modelo y lógica relacionada con gastos
├── leerCSV/
│   └── leerCSV.go          # Funciones para leer y procesar archivos CSV
├── main.go                 # Punto de entrada del programa
├── go.mod                  # Módulo Go
└── go.sum                  # Dependencias
```

## Características

- **Base de datos SQLite:**
  - Crea y administra la tabla `Gasto`.
  - Almacena información como descripción, monto, fecha y categoría.
  - Proporciona funciones para consultar registros existentes.

- **Lectura de Archivos CSV:**
  - Procesa un archivo CSV (`activity.csv`) con datos financieros.
  - Valida si es un formato AMEX y extrae información relevante.

- **Gestión de Gastos:**
  - Genera registros de gastos a partir de entradas del archivo CSV.
  - Modelo estructurado con soporte para extensiones futuras.

## Dependencias

- **SQLite Driver:**
  - `github.com/mattn/go-sqlite3` para interactuar con SQLite.

Instala las dependencias con:

```bash
go mod tidy
```

## Cómo Ejecutar

1. **Asegúrate de tener Go instalado.**
2. **Clona el repositorio.**
3. **Ejecuta el programa:**

   ```bash
   go run main.go
   ```

4. El programa:
   - Inicializará la base de datos.
   - Consultará los registros existentes.
   - Procesará el archivo `activity.csv`.

## Formato del CSV

El archivo `activity.csv` debe tener las siguientes columnas (o similares, dependiendo de la función `leerCSV`):

| Fecha       | ... | Descripción       | Titular       | ... | Importe  |
|-------------|-----|-------------------|---------------|-----|----------|
| 2024-11-01  | ... | Supermercado      | Arturo Zepeda | ... | 125.50   |

## Contribuciones

1. Haz un fork del repositorio.
2. Crea un branch para tus cambios:
   ```bash
   git checkout -b feature/nueva-funcionalidad
   ```
3. Haz un pull request.

## Licencia

Este proyecto está bajo la licencia [MIT](LICENSE). Siéntete libre de usarlo, modificarlo y compartirlo.

