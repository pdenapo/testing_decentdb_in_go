package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/sphildreth/decentdb/bindings/go/decentdb-go"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

/* Example form  bun orm documentation
https://bun.uptrace.dev/guide/golang-orm.html#defining-models */
/* autoincrement is not supported */

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64     `bun:",pk"`
	Name      string    `bun:",notnull"`
	Email     string    `bun:",unique,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func main() {
	print("testing bun orm for decentdb")
	sqldb, err := sql.Open("decentdb", "file:/tmp/app.ddb")
	if err != nil {
		log.Fatal(err)
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	ctx := context.Background()

	_, err = db.NewCreateTable().
		Model((*User)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`INSERT INTO "users" ("id", "name", "email", "created_at", "updated_at") VALUES (0, 'Alice', 'alice@example.com', DEFAULT, DEFAULT), (0, 'Bob', 'bob@example.com', DEFAULT, DEFAULT) RETURNING "created_at", "updated_at"INSERT INTO "users" ("id", "name", "email", "created_at", "updated_at") VALUES (0, 'Alice', 'alice@example.com', DEFAULT, DEFAULT), (0, 'Bob', 'bob@example.com', DEFAULT, DEFAULT) RETURNING "created_at", "updated_at"`)
	if err != nil {
		log.Fatal(err)
	}

	users := []*User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}

	_, err = db.NewInsert().Model(&users).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
