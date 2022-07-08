package main

import (
	"fmt"
	"go_sample_injection/config"
	"go_sample_injection/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()

	db := cfg.DbConn()

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)


	routeEngine := gin.Default()
	routerGroup := routeEngine.Group("/api")
	routerGroup.POST("/auth/login", func(ctx *gin.Context) {
		var login model.Login

		if err := ctx.ShouldBindJSON(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error(),
			})
		}

		var useCred = model.UserCredential{}

		// sql := fmt.Sprintf("SELECT * FROM user_credential WHERE user_name='%s' and user_passwords='%s' and is_blocked='f'", login.User, login.Password)
		//WHAT=WHAT --> TRUE, RETIURN ALL ROWS
		//gak aman 

		sql := "SELECT * FROM user_credential WHERE user_name=$1 and user_password=$2 and is_blocked='f'"
		//lebih aman
		//solusi nya agar tidak terkena sql injetion menggunakan query param seperto $1, dan atau ?

		//pake query param seperto biar hasil what=what nya gak masuk ke querry

		//sanitizai namanya

		//where itu balikin booolean makanya ketika terima true dia lgsng balik all



		log.Println("sql: ", sql)

		err := db.Get(&useCred, sql)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return 
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "PONG",
		})


	})

	var apiPort = config.ApiConfig{}
	apiPort.ApiHost = os.Getenv("API_HOST")
	apiPort.ApiPort = os.Getenv("API_PORT")

	listenAddress := fmt.Sprintf("%s:%s", apiPort.ApiHost, apiPort.ApiPort)

	err := routeEngine.Run(listenAddress)

	if err != nil {
		panic(err)
	}
}

