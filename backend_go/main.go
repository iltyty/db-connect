package main

import (
	"github.com/iltyty/db_connect/backend_go/models"
	"github.com/iltyty/db_connect/backend_go/routers/api/v1"
)

func main() {
	models.Init()
	_ = v1.Init()
}
