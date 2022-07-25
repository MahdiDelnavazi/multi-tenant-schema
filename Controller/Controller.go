package Controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"multi-tenant-postgres/Database"
)

type Controller struct {
	db Database.Database
}

func NewUserController(db Database.Database) *Controller {
	return &Controller{db: db}
}

func (controller *Controller) MigrateUp(context *gin.Context) {
	tenantName := context.GetHeader("tenant-name")

	err := controller.db.CheckSchemaExist(tenantName)
	if err != nil {
		log.Println("schema exist : ", err)
		return
	}

	err = controller.db.CreateNewTenant(tenantName)
	if err != nil {
		log.Println("query error create new tenant : ", err)
		return
	}

	err = controller.db.CreateTableA(tenantName)
	if err != nil {
		log.Println("query error create table A : ", err)
		return
	}

	err = controller.db.CreateTableB(tenantName)
	if err != nil {
		log.Println("query error create table B : ", err)
		return
	}

	err = controller.db.CreateTableC(tenantName)
	if err != nil {
		log.Println("query error create table C : ", err)
		return
	}

	err = controller.db.CreateTableD(tenantName)
	if err != nil {
		log.Println("query error create table D : ", err)
		return
	}

	err = controller.db.CreateTableE(tenantName)
	if err != nil {
		log.Println("query error create table E : ", err)
		return
	}

	err = controller.db.CreateTableF(tenantName)
	if err != nil {
		log.Println("query error create table F : ", err)
		return
	}

	err = controller.db.CreateTableG(tenantName)
	if err != nil {
		log.Println("query error create table G : ", err)
		return
	}

	err = controller.db.CreateTableH(tenantName)
	if err != nil {
		log.Println("query error create table H : ", err)
		return
	}

	err = controller.db.CreateTableI(tenantName)
	if err != nil {
		log.Println("query error create table I : ", err)
		return
	}
}
