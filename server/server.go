package main

import (
	"database/sql"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"server/controller"
)

func main() {
	if _,err := os.Stat("/kouji"); os.IsNotExist(err){
		os.Mkdir("/kouji",0777)
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	db, err := sql.Open("mysql", "root:root@([db]:3306)/NUPS")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userctr := &controller.NupsCtr{DB: db}

	router.Use(cors.New(config))
	router.GET("/", userctr.ListKi)
	router.GET("/:ki", userctr.ListKouji)
	router.POST("/:ki", userctr.AddKi)
	router.GET("/:ki/:id")
	router.POST("/:ki/:id", userctr.UploadKouji)
	router.PATCH("/:ki/:id", userctr.UpdateKouji)
	router.DELETE("/:ki/:id", userctr.DeleteKouji)
	router.Run(":8888")
}
