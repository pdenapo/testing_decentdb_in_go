use decentdb::{Db, DbConfig, DbError, Value};

// ─── SQL y datos compartidos ───────────────────────────────────────────────

const CREATE_USERS: &str = r#"
CREATE TABLE IF NOT EXISTS users (
    id   INTEGER NOT NULL,
    name TEXT    NOT NULL
)"#;

const INSERT_USERS: &str = r#"
INSERT INTO users (id, name) VALUES (1, 'Alice'), (2, 'Bob')
"#;

const SELECT_USERS: &str = r#"SELECT * FROM users ORDER BY id"#;
const COUNT_USERS: &str = r#"SELECT COUNT(*) FROM users"#;

/// Datos que ambos backends deben devolver.
fn expected_users() -> Vec<(i64, String)> {
    vec![(1, "Alice".to_string()), (2, "Bob".to_string())]
}

// ─── Test DecentDB ─────────────────────────────────────────────────────────

#[cfg(test)]
mod tests_decentdb {
    use super::*;

    #[test]
    fn test_basic() -> Result<(), DbError> {
        let db = Db::open(":memory:", DbConfig::default())?;

        db.execute(CREATE_USERS)?;
        db.execute(INSERT_USERS)?;

        // SELECT *
        let result = db.execute(SELECT_USERS)?;
        assert_eq!(result.rows().len(), 2);
        assert_eq!(result.columns(), &["id", "name"]);

        for (i, (id, name)) in expected_users().into_iter().enumerate() {
            assert_eq!(
                result.rows()[i].values(),
                &[Value::Int64(id), Value::Text(name)]
            );
        }

        // COUNT(*)
        let result_count = db.execute(COUNT_USERS)?;
        assert_eq!(result_count.rows().len(), 1);
        assert_eq!(result_count.rows()[0].values(), &[Value::Int64(2)]);

        Ok(())
    }
}

// ─── Test Turso (crate `turso`) ────────────────────────────────────────────

#[cfg(test)]
mod tests_turso {
    use super::*;
    use turso::Builder;

    #[tokio::test]
    async fn test_basic() -> Result<(), Box<dyn std::error::Error>> {
        let db = Builder::new_local(":memory:").build().await?;
        let conn = db.connect()?;

        conn.execute(CREATE_USERS, ()).await?;
        conn.execute(INSERT_USERS, ()).await?;

        // SELECT *
        let mut rows = conn.query(SELECT_USERS, ()).await?;
        let mut got = Vec::new();
        while let Some(row) = rows.next().await? {
            let id: i64 = row.get(0)?;
            let name: String = row.get(1)?;
            got.push((id, name));
        }
        assert_eq!(got, expected_users());

        // COUNT(*)
        let mut rows = conn.query(COUNT_USERS, ()).await?;
        let row = rows.next().await?.expect("una fila de COUNT");
        let count: i64 = row.get(0)?;
        assert_eq!(count, 2);

        Ok(())
    }
}

fn main() {
    println!("This is a test program. Use cargo test.");
}
