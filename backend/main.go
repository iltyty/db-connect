package main

import (
	"github.com/iltyty/db_connect/backend_go/conf"
	"github.com/iltyty/db_connect/backend_go/models"
	"github.com/iltyty/db_connect/backend_go/routers/api/v1"
	"github.com/iltyty/db_connect/backend_go/utils"
)

func main() {
	conf.Init()
	models.Init()
	utils.InitRedis()
	_ = v1.Init()
}
