package main

import (
	"flag"
	"log"

	"gitlab.com/nguyencatpham/go-effective-study/cmd/migration"
	"gitlab.com/nguyencatpham/go-effective-study/pkg/api"

	"gitlab.com/nguyencatpham/go-effective-study/pkg/utl/config"
)

func main() {

	cfgPath := flag.String("p", "./cmd/api/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)
	checkErr(migration.Init(cfg))
	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
