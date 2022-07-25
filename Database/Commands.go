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

// CreateUUID for create new user table in tenant
func (db *Database) CreateUUID() error {
	query := fmt.Sprintf(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
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

// CreateTableA for create new user table in tenant
func (db *Database) CreateTableA(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."A"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()));`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableB for create new user table in tenant
func (db *Database) CreateTableB(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."B"("Id" serial PRIMARY KEY,
		"IdA" int,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		CONSTRAINT fk_B
		FOREIGN KEY(IdA) 
	  	REFERENCES A(Id));`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableC for create new user table in tenant
func (db *Database) CreateTableC(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."C"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()));`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableD for create new user table in tenant
func (db *Database) CreateTableD(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."D"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableE for create new user table in tenant
func (db *Database) CreateTableE(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."E"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Name" string);`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableF for create new user table in tenant
func (db *Database) CreateTableF(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."F"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()));`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableG for create new user table in tenant
func (db *Database) CreateTableG(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."G"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()));`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableH for create new user table in tenant
func (db *Database) CreateTableH(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."H"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()));`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableI for create new user table in tenant
func (db *Database) CreateTableI(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."I"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()));`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableJ for create new user table in tenant
func (db *Database) CreateTableJ(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."J"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()));`, tenant)
	_, err := db.NameSpace.Exec(query)
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
