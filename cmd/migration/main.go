package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"gitlab.com/nguyencatpham/gorsk/pkg/utl/config"
)

func main() {
	ENV := os.Getenv("ENV")
	if len(ENV) == 0 {
		ENV = "dev"
	}
	cfg, err := config.Load(ENV)
	if err == nil && cfg.DB != nil {
		database := cfg.DB.DBName
		isMigrate, err := createDatabase(database, cfg.DB.PSNBase)
		if isMigrate && err == nil {
			log.Println("Migration is starting...")
			const sqlFuncs = `
			CREATE OR REPLACE FUNCTION f_concat_ws(text, VARIADIC text[])
  			RETURNS text LANGUAGE sql IMMUTABLE AS 'SELECT array_to_string($2, $1)';
			`

			u, err := pg.ParseURL(cfg.DB.PSN)
			checkErr(err)
			db := pg.Connect(u)
			defer db.Close()

			_, err = db.Exec("SELECT 1")
			checkErr(err)
			///create shard
			_, err = db.Exec(sqlFuncs)
			checkErr(err)

			createSchema(db)
		}
	}
}
func createDatabase(name string, psnBase string) (bool, error) {
	db, err := sql.Open("postgres", psnBase)
	defer db.Close()
	if err != nil {
		return false, err
	}
	query := fmt.Sprintf(`SELECT 1 FROM pg_database WHERE datname = '%s';`, name)
	result, err := db.Exec(query)
	if err != nil {
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows == 0 {
		log.Printf(`Database "%s" is creating...`, name)
		query := fmt.Sprintf(`CREATE DATABASE %s;`, name)
		_, err := db.Exec(query)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createSchema(db *pg.DB, models ...interface{}) {
	for _, model := range models {
		checkErr(db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		}))
	}
}
