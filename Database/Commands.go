package Database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

// CreateNewTenant for create new tenant
func (db *Database) CreateNewTenant(tenant string) error {
	query := fmt.Sprintf(`create schema %s`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CheckSchemaExist check if schema exist or not
func (db *Database) CheckSchemaExist(tenant string) error {
	query := fmt.Sprintf(`SELECT schema_name FROM information_schema.schemata WHERE schema_name = %s;`, tenant)
	result, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	if result != nil {
		// if rows == 0 it means schema is not found , but row == 1 means schema exists
		rows, _ := result.RowsAffected()
		if rows == 0 {
			return fmt.Errorf("schema not found %s", tenant)
		}
	}
	return nil
}

// CreateUserTable for create new user table in tenant
func (db *Database) CreateUserTable(tenant string) error {
	_, err := db.NameSpace.Exec(`create table if not exists $1."User"();`, tenant)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) MigrateUp() {
	driver, err := postgres.WithInstance(db.NameSpace, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"../Database/Schema",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	m.Up()
}
