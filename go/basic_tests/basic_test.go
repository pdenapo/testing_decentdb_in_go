package bun_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/sphildreth/decentdb/bindings/go/decentdb-go"
	_ "turso.tech/database/tursogo"
)

func testInsert(t *testing.T, driver string) {
	t.Log("Runign test for driver ", driver)
	db, err := sql.Open(driver, ":memory:")
	if err != nil {
		t.Fatal("Error en Open ", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "users" ("id" INT NOT NULL, "name" VARCHAR NOT NULL)`)
	if err != nil {
		t.Fatal("Error en Create ", err)

	}

	_, err = db.Exec(`INSERT INTO "users" ("id", "name") VALUES (1, 'Alice'), (2, 'Bob')`)
	if err != nil {
		t.Fatal("Error en Insert", err)
	}

	// 1. Usar db.Query en lugar de db.Exec para un SELECT
	rows, err := db.Query(`SELECT id, name FROM users ORDER BY id`)
	if err != nil {
		t.Fatal("Error en Select", err)
	}
	defer rows.Close() // Es importante cerrar las filas al terminar

	// 2. Iterar sobre cada fila devuelta
	for rows.Next() {
		var id int
		var name string

		// 3. Escanear los valores de la fila actual a las variables
		if err := rows.Scan(&id, &name); err != nil {
			t.Fatal("Error en Scan", err)
		}

		// 4. Mostrar el resultado (t.Logf imprime en la consola si hay fallo o si usas go test -v)
		t.Logf("Usuario: id = %d, name = %s", id, name)

		switch id {
		case 1:
			if name != "Alice" {
				t.Error("Se esperaba Alice")
			}
		case 2:
			if name != "Bob" {
				t.Error("Se esperaba Bob")
			}
		}
	}

	// 5. Verificar si hubo algún error durante la iteración
	if err := rows.Err(); err != nil {
		t.Fatal("Error en la iteración de filas", err)
	}
	// 1. Usar db.QueryRow para consultas que devuelven un único resultado
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count)
	if err != nil {
		t.Fatal("Error en Count", err)
	}

	// 2. Mostrar o verificar el resultado
	t.Logf("Total de usuarios: %d", count)

	// Opcional: Puedes hacer una aserción si estás probando un valor esperado
	if count != 2 {
		t.Fatalf("Se esperaban 2 usuarios, se obtuvieron %d", count)
	}
}

func TestInsertSqlite(t *testing.T) {
	testInsert(t, "sqlite3")
}

func TestInsertDecentDb(t *testing.T) {
	testInsert(t, "decentdb")
}

func TestInsertDecentTurso(t *testing.T) {
	testInsert(t, "turso")
}
