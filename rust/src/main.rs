#[cfg(test)]
mod tests {
    //use super::*;
    use decentdb::{Db, DbConfig, DbError, Value};

    #[test]
    fn test_decentdb_basic() -> Result<(), DbError> {
        let db = Db::open(":memory:", DbConfig::default())?;

        db.execute(
            r#"CREATE TABLE IF NOT EXISTS "users" ("id" INT NOT NULL, "name" VARCHAR NOT NULL)"#,
        )?;
        db.execute(r#"INSERT INTO "users" ("id", "name") VALUES (1, 'Alice'), (2, 'Bob')"#)?;

        let result = db.execute(r#"SELECT * FROM "users" ORDER BY id"#)?;

        // 1. Comprobar cantidad de filas
        assert_eq!(result.rows().len(), 2);

        // 2. Comprobar nombres de columnas
        assert_eq!(result.columns(), &["id", "name"]);

        // 3. Comprobar el contenido de cada fila
        assert_eq!(
            result.rows()[0].values(),
            &[Value::Int64(1), Value::Text("Alice".to_string())]
        );
        assert_eq!(
            result.rows()[1].values(),
            &[Value::Int64(2), Value::Text("Bob".to_string())]
        );

        let result_count = db.execute(r#"SELECT COUNT(*) FROM "users""#)?;
        assert_eq!(result_count.rows().len(), 1);
        assert_eq!(result_count.rows()[0].values(), &[Value::Int64(2)]);
        Ok(())
    }
}

fn main() {
    println!("This is a test program.Use cargo test.")
}
