package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	env := os.Getenv("DB_CONNECT_STRING")
	fmt.Println(env)
	db, err := gorm.Open(mysql.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	ogDB, err := db.DB()
	if err != nil {
		log.Fatal()
	}
	defer ogDB.Close()

	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		var numb int
		err := db.Raw("Select 2").Find(&numb).Error
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, map[string]any{
			"message": "ok",
			"data":    numb,
		})
	})

	err = router.Run(":8082")
	if err != nil {
		log.Fatal(err)
	}
}
