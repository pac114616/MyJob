package main

import (

	"github.com/gin-gonic/gin" //引入gin
	_ "github.com/go-sql-driver/mysql"
	"text/common"
)

func main() {
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}


