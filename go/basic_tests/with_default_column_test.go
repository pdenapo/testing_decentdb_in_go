package basic_tests

import (
	"database/sql"
	"testing"

	_ "github.com/duckdb/duckdb-go/v2"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sphildreth/decentdb/bindings/go/decentdb-go"
	_ "turso.tech/database/tursogo"
)

func testInsertWithDefaultColumn(t *testing.T, driver string) {
	t.Log("Runign test for driver ", driver)
	db, err := sql.Open(driver, ":memory:")
	if err != nil {
		t.Fatal("Error en Open ", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE t (id INT PRIMARY KEY, x INT DEFAULT 42);`)
	if err != nil {
		t.Fatal("Error en Create ", err)

	}

	_, err = db.Exec(`INSERT INTO t (id, x) VALUES (1, DEFAULT);`)
	if err != nil {
		t.Fatal("Error en Insert", err)
	}

	// 1. Usar db.Query en lugar de db.Exec para un SELECT
	rows, err := db.Query(`SELECT * FROM t`)
	if err != nil {
		t.Fatal("Error en Select", err)
	}
	defer rows.Close() // Es importante cerrar las filas al terminar

	// 2. Iterar sobre cada fila devuelta
	for rows.Next() {
		var id int
		var x int

		// 3. Escanear los valores de la fila actual a las variables
		if err := rows.Scan(&id, &x); err != nil {
			t.Fatal("Error en Scan", err)
		}

		// 4. Mostrar el resultado (t.Logf imprime en la consola si hay fallo o si usas go test -v)
		t.Logf("Usuario: id = %d, name = %d", id, x)

		switch id {
		case 1:
			if x != 42 {
				t.Error("Se esperaba 42")
			}
		}
	}

	// 5. Verificar si hubo algún error durante la iteración
	if err := rows.Err(); err != nil {
		t.Fatal("Error en la iteración de filas", err)
	}
	// 1. Usar db.QueryRow para consultas que devuelven un único resultado
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM t`).Scan(&count)
	if err != nil {
		t.Fatal("Error en Count", err)
	}

	// 2. Mostrar o verificar el resultado
	t.Logf("Total de usuarios: %d", count)

	// Opcional: Puedes hacer una aserción si estás probando un valor esperado
	if count != 1 {
		t.Fatalf("Se esperaban 1 fila, se obtuvieron %d", count)
	}
}

/* sqlite,turso do not support this */

func TestInsertWithDefaultColumDuckDb(t *testing.T) {
	testInsertWithDefaultColumn(t, "duckdb")
}

func TestInsertWithDefaultColumDecentdb(t *testing.T) {
	testInsertWithDefaultColumn(t, "decentdb")
}
