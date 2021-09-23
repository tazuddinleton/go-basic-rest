package main

import (
	"github.com/tazuddinleton/go/basicrest/conf"
	"github.com/tazuddinleton/go/basicrest/db"
	"github.com/tazuddinleton/go/basicrest/logger"
)

func main() {
	c := conf.New()
	file := logger.Configure(c)
	defer file.Close()
	db := db.New(c.Db())
	defer db.Close()

}
