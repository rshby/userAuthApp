package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	db "userAuthApp/db/connection"
	"userAuthApp/handler"
	"userAuthApp/repository"
	"userAuthApp/router"
	"userAuthApp/service"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error cant load env file : " + err.Error())
	}

	fmt.Println(os.Getenv("DB_NAME"))
}

func main() {
	db := db.ConnecDB()

	// register repository layer
	accountRepo := repository.NewAccountRepository(db)

	// register service layer
	accountService := service.NewAccountService(db, accountRepo)

	// register handler layer
	accountHandler := handler.NewAccountHandler(accountService)

	r := gin.Default()
	apiRoute := r.Group("/api/v1")

	// create account routes
	router.CreateAccountRoutes(apiRoute, accountHandler)

	appPort := os.Getenv("APP_PORT")
	r.Run(":" + appPort)
}
