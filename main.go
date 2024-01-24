package main

import (
	"log"
	"rest-go-demo/controllers"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := gin.Default()
	initalizeHandlers(router)
	log.Fatal(router.Run(":8090"))
}

func initalizeHandlers(router *gin.Engine) {
	router.POST("/create", controllers.CreatePerson)
	router.GET("/get", controllers.GetAllPerson)
	router.GET("/get/:id", controllers.GetPersonByID)
	router.PUT("/update/:id", controllers.UpdatePersonByID)
	router.DELETE("/delete/:id", controllers.DeletePersonByID)
}

func initDB() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "password",
		DB:         "learning",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Person{})
}
