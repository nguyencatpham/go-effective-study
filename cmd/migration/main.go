package migration

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/config"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/helper"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/model"
)

// Init func
func Init(cfg *config.Configuration) error {
	if cfg.DB == nil {
		return helper.HandleError("Missing config database")
	}
	updateSwagger(cfg)
	isMigrate, err := createDatabase(cfg.DB.DBName, cfg.DB.PSNBase)
	log.Println("open connection ", isMigrate)
	if !isMigrate {
		return helper.HandleError("No need migrate. Starting program...")
	}
	if err != nil {
		return err
	}
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

	createSchema(db, &model.Topic{}, &model.Role{}, &model.User{}, &model.TopicDetail{}, &model.UserDetail{}, &model.Company{}, &model.Location{})
	seedData(db)
	return nil
}
func createDatabase(name string, psnBase string) (bool, error) {
	db, err := sql.Open("postgres", psnBase)
	defer db.Close()
	if err != nil {
		log.Println("open connection error")
		return false, err
	}
	query := fmt.Sprintf(`SELECT 1 FROM pg_database WHERE datname = '%s';`, name)
	result, err := db.Exec(query)
	if err != nil {
		log.Println("exec db error", query, err)
		return false, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("get result error")
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
func seedData(db *pg.DB) {
	dbInsert := `INSERT INTO public.companies VALUES (1, now(), now(), NULL, 'admin_company', true);
	INSERT INTO public.locations VALUES (1, now(), now(), NULL, 'admin_location', true, 'admin_address', 1);
	INSERT INTO public.roles VALUES (100, 100, 'SUPER_ADMIN');
	INSERT INTO public.roles VALUES (110, 110, 'ADMIN');
	INSERT INTO public.roles VALUES (120, 120, 'COMPANY_ADMIN');
	INSERT INTO public.roles VALUES (130, 130, 'LOCATION_ADMIN');
	INSERT INTO public.roles VALUES (200, 200, 'USER');`
	queries := strings.Split(dbInsert, ";")
	for _, v := range queries[0 : len(queries)-1] {
		_, err := db.Exec(v)
		checkErr(err)
	}
}
func updateSwagger(cfg *config.Configuration) {
	//-- replace host swagger
	if cfg.Server.Host != "" {
		swaggerJSON, err := helper.Readfile(cfg.Server.SwaggerJSON)
		checkErr(err)

		m := make(map[string]interface{})
		err = json.Unmarshal([]byte(swaggerJSON), &m)
		checkErr(err)
		value, _ := m["host"].(string)
		swaggerJSON = strings.Replace(swaggerJSON, value, cfg.Server.Host, -1)
		swaggerJSON = strings.Replace(swaggerJSON, "http", "https", -1)
		err = helper.WriteFile(cfg.Server.SwaggerJSON, swaggerJSON)
		checkErr(err)
		log.Println("update swagger host...DONE")
	}
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
