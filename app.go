package main

import (
	"fmt"

	"github.com/tazuddinleton/go/basicrest/conf"
	"github.com/tazuddinleton/go/basicrest/logger"
)

func main() {
	c := conf.New()
	file := logger.Configure(c)

	defer file.Close()

	fmt.Println(*c)
}
