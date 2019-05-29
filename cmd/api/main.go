package main

import (
	"flag"
	"log"

	"github.com/nguyencatpham/go-effective-study/cmd/migration"
	"github.com/nguyencatpham/go-effective-study/pkg/api"

	"github.com/nguyencatpham/go-effective-study/pkg/utl/config"
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
