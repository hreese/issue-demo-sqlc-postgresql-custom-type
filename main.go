//go:generate go tool github.com/sqlc-dev/sqlc/cmd/sqlc generate
package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/hreese/issue-demo-sqlc-postgresql-custom-type/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var ctx = context.Background()

	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// run this for every connection in the connection pool after the database connection is established
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// create type registration definition (both scalar and array)
		typeNames := []string{"MyEnum", "_MyEnum"}
		types, err := conn.LoadTypes(ctx, typeNames)
		if err != nil {
			return err
		}
		// register custom types
		conn.TypeMap().RegisterTypes(types)
		log.Printf("Registered %s types", strings.Join(typeNames, ", "))

		// https://github.com/jackc/pgx/issues/953 ?
		//conn.TypeMap().RegisterDefaultPgType(database.Myenum("one"), `MyEnum`)
		//conn.TypeMap().RegisterDefaultPgType([]database.Myenum{"one"}, `_MyEnum`)
		return nil
	}

	// start up db pool
	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	// create record with custom type
	data := []database.Myenum{
		database.MyenumOne,
		database.MyenumTwo,
		database.MyenumThree,
		database.MyenumFour,
	}
	// insert record with custom type
	q := database.New(dbPool)
	err = q.InsertNewEntry(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
}
