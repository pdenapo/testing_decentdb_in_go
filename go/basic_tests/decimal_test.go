package basic_tests

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sphildreth/decentdb/bindings/go/decentdb-go"
	_ "turso.tech/database/tursogo"
)

func testDecimal(t *testing.T, driver string, expected bool) {
	t.Log("Runing test for driver ", driver)
	db, err := sql.Open(driver, ":memory:")
	if err != nil {
		t.Fatal("Error en Open ", err)
	}
	defer db.Close()

	// 1. Usar db.Query en lugar de db.Exec para un SELECT
	rows, err := db.Query(`select cast(0.1 as numeric)+cast(0.2 as numeric)=cast(0.3 as numeric)`)
	if err != nil {
		t.Fatal("Error en Select", err)
	}
	defer rows.Close() // Es importante cerrar las filas al terminar

	// 2. Iterar sobre cada fila devuelta
	for rows.Next() {
		var x bool

		// 3. Escanear los valores de la fila actual a las variables
		if err := rows.Scan(&x); err != nil {
			t.Fatal("Error en Scan", err)
		}

		// 4. Mostrar el resultado (t.Logf imprime en la consola si hay fallo o si usas go test -v)
		t.Logf("x=%v", x)

		if x != expected {
			t.Fail()
		}
	}

	// 5. Verificar si hubo algún error durante la iteración
	if err := rows.Err(); err != nil {
		t.Fatal("Error en la iteración de filas", err)
	}
}

// Duckdb and Decentdb have exact decimals, so the comparison is true.
// Sqlite and turso use float, so the comparison is false

func TestDecimalqlite(t *testing.T) {
	testDecimal(t, "sqlite3", false)
}

func TestDecimalDecentDb(t *testing.T) {
	testDecimal(t, "decentdb", true)
}

func TestDecimalTurso(t *testing.T) {
	testDecimal(t, "turso", false)
}

func TestDecimalDuckDb(t *testing.T) {
	testDecimal(t, "duckdb", true)
}
