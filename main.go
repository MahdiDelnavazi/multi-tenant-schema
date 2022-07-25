package main

import (
	"github.com/gin-gonic/gin"
	"multi-tenant-postgres/Controller"
	"multi-tenant-postgres/Database"
)

func main() {

	dbConfig := Database.Config{Database: "multi-tenant-schema", Host: "localhost", Port: "5432", User: "root", Password: "secret"}
	database, err := Database.New(dbConfig)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	controller := Controller.NewUserController(database)
	router := gin.Default()
	router.PUT("/create-order/", controller.MigrateUp)
	router.Run()

}
