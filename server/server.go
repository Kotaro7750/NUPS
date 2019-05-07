package main

import (
	"database/sql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"server/controller"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	db, err := sql.Open("")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userctr := &controller.UserCtr{DB: db}

	router.Use(cors.New(config))
	router.GET("/",ListKi)
	router.GET("/:ki",ListKouji)
	router.POST("/:ki",AddKi)
	router.GET("/:ki/:id",)
	router.POST("/:ki/:id",UploadKouji)
	router.UPDATE("/:ki/:id",UpdateKouji)
	router.DELETE("/:ki/:id",DeleteKouji)
	router.Run(":8080")
}
