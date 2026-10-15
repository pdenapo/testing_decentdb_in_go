package basic_tests

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sphildreth/decentdb/bindings/go/decentdb-go"
	_ "turso.tech/database/tursogo"
)

func testPgTypeOf(t *testing.T, driver string, expected string) {
	t.Log("Runing test for driver ", driver)
	db, err := sql.Open(driver, ":memory:")
	if err != nil {
		t.Fatal("Error en Open ", err)
	}
	defer db.Close()

	// 1. Usar db.Query en lugar de db.Exec para un SELECT
	rows, err := db.Query(`SELECT pg_typeof(to_timestamp(1711929600123456 / 1000000.0))`)
	if err != nil {
		t.Fatal("Error en Select", err)
	}
	defer rows.Close() // Es importante cerrar las filas al terminar

	// 2. Iterar sobre cada fila devuelta
	for rows.Next() {
		var x string

		// 3. Escanear los valores de la fila actual a las variables
		if err := rows.Scan(&x); err != nil {
			t.Fatal("Error en Scan", err)
		}

		// 4. Mostrar el resultado (t.Logf imprime en la consola si hay fallo o si usas go test -v)
		t.Logf("x=%s", x)

		if x != expected {
			t.Fail()
		}
	}

	// 5. Verificar si hubo algún error durante la iteración
	if err := rows.Err(); err != nil {
		t.Fatal("Error en la iteración de filas", err)
	}
}

func TestPgTypeOfDecentDb(t *testing.T) {
	testPgTypeOf(t, "decentdb", "timestamp with time zone")
}

func TestPgTypeOfDuckDb(t *testing.T) {
	testPgTypeOf(t, "duckdb", "timestamp with time zone")
}
