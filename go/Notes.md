Notes
=====

Known differences from postgresql affecting bun

autoincrement =>

2026/10/13 20:58:37 decentdb error 5: SQL error: type bigserial is not supported

[bun]  10:26:03.163   INSERT                1.096ms  INSERT INTO "users" ("id", "name", "email", "created_at", "updated_at") VALUES (0, 'Alice', 'alice@example.com', DEFAULT, DEFAULT), (0, 'Bob', 'bob@example.com', DEFAULT, DEFAULT) RETURNING "created_at", "updated_at"INSERT INTO "users" ("id", "name", "email", "created_at", "updated_at") VALUES (0, 'Alice', 'alice@example.com', DEFAULT, DEFAULT), (0, 'Bob', 'bob@example.com', DEFAULT, DEFAULT) RETURNING "created_at", "updated_at" 	  *decentdb.DecentDBError: decentdb error 5: SQL error: Error splitting: syntax error at or near "INTO" 

[bun]  10:26:03.164   INSERT                1.751ms  INSERT INTO "users" ("id", "name", "email", "created_at", "updated_at") VALUES (0, 'Alice', 'alice@example.com', DEFAULT, DEFAULT), (0, 'Bob', 'bob@example.com', DEFAULT, DEFAULT) RETURNING "created_at", "updated_at" 	  *decentdb.DecentDBError: decentdb error 5: SQL error: expression node Node is not supported
