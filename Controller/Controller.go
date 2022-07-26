package Controller

import (
	"fmt"
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

func (controller *Controller) SchemaMigrateUp(context *gin.Context) {
	//tenantName := context.GetHeader("tenant-name")
	for i := 0; i <= 10; i++ {
		log.Println("create tenant")
		tenantName := fmt.Sprintf("tenant%d", i)

		err := controller.db.CreateNewSchemaTenant(tenantName)
		if err != nil {
			log.Println("query error create new tenant : ", err)
			return
		}

		err = controller.db.CreateTableA(tenantName)
		if err != nil {
			log.Println("query error create table A : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableA(tenantName)
			if err != nil {
				log.Println("query error insert into table A : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableA(tenantName)
		if err != nil {
			log.Println("query error index table A : ", err)
			return
		}

		err = controller.db.CreateTableB(tenantName)
		if err != nil {
			log.Println("query error create table B : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableB(tenantName)
			if err != nil {
				log.Println("query error insert into table B : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableB(tenantName)
		if err != nil {
			log.Println("query error index table B : ", err)
			return
		}

		err = controller.db.CreateTableC(tenantName)
		if err != nil {
			log.Println("query error create table C : ", err)
			return
		}

		for i := 0; i <= 150000; i++ {
			err = controller.db.InsertIntoTableC(tenantName)
			if err != nil {
				log.Println("query error insert into table C : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableC(tenantName)
		if err != nil {
			log.Println("query error index table C : ", err)
			return
		}

		err = controller.db.CreateTableD(tenantName)
		if err != nil {
			log.Println("query error create table D : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableD(tenantName)
			if err != nil {
				log.Println("query error insert into table D : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableD(tenantName)
		if err != nil {
			log.Println("query error index table D : ", err)
			return
		}

		err = controller.db.CreateTableE(tenantName)
		if err != nil {
			log.Println("query error create table E : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableE(tenantName)
			if err != nil {
				log.Println("query error insert into table E : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableE(tenantName)
		if err != nil {
			log.Println("query error index table E : ", err)
			return
		}

		err = controller.db.CreateTableF(tenantName)
		if err != nil {
			log.Println("query error create table F : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableF(tenantName)
			if err != nil {
				log.Println("query error insert into table F : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableF(tenantName)
		if err != nil {
			log.Println("query error index table F : ", err)
			return
		}

		err = controller.db.CreateTableG(tenantName)
		if err != nil {
			log.Println("query error create table G : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableG(tenantName)
			if err != nil {
				log.Println("query error insert into table G : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableG(tenantName)
		if err != nil {
			log.Println("query error index table G : ", err)
			return
		}

		err = controller.db.CreateTableH(tenantName)
		if err != nil {
			log.Println("query error create table H : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableH(tenantName)
			if err != nil {
				log.Println("query error insert into table H : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableH(tenantName)
		if err != nil {
			log.Println("query error index table H : ", err)
			return
		}

		err = controller.db.CreateTableI(tenantName)
		if err != nil {
			log.Println("query error create table I : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableI(tenantName)
			if err != nil {
				log.Println("query error insert into table I : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableI(tenantName)
		if err != nil {
			log.Println("query error index table I : ", err)
			return
		}

		err = controller.db.CreateTableJ(tenantName)
		if err != nil {
			log.Println("query error create table J : ", err)
			return
		}

		for i := 0; i <= 20; i++ {
			err = controller.db.InsertIntoTableJ(tenantName)
			if err != nil {
				log.Println("query error insert into table J : ", err)
				return
			}
		}

		err = controller.db.CreateIndexTableJ(tenantName)
		if err != nil {
			log.Println("query error index table I : ", err)
			return
		}

	}
}
