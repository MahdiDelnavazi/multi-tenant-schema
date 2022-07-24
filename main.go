package main

import (
	"github.com/gin-gonic/gin"
	"multi-tenant-postgres/Database"
)

func main() {

	dbConfig := Database.Config{Database: "multi-tenant-schema", Host: "localhost", Port: "5432", User: "root", Password: "secret"}
	Database.New(dbConfig)
	router := gin.Default()
	//router.PUT("/create-order/", Repository.CreateOrder)
	router.Run()

}
