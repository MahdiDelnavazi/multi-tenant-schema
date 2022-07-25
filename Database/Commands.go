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
	query := fmt.Sprintf(`create schema if not exists %s`, tenant)
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
	query := fmt.Sprintf(`SELECT schema_name FROM information_schema.schemata WHERE schema_name = '%s';`, tenant)
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
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableA for insert seed data into table A
func (db *Database) InsertIntoTableA(tenant string) error {
	query := fmt.Sprintf(`insert into %s."A" ("Age") values (%d);`, tenant, 12)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableA for create multi column index in table A
func (db *Database) CreateIndexTableA(tenant string) error {
	query := fmt.Sprintf(`create index A on %s."A"("Id","CreatedAt","Age")`, tenant)

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
		"Age" int,
		CONSTRAINT fk_B
		FOREIGN KEY("IdA") 
	  	REFERENCES %s."A"("Id"));`, tenant, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableB for insert seed data into table B
func (db *Database) InsertIntoTableB(tenant string) error {
	query := fmt.Sprintf(`insert into %s."B" ("Age","IdA") values (%d,%d);`, tenant, 12, 1)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableB for create multi column index in table B
func (db *Database) CreateIndexTableB(tenant string) error {
	query := fmt.Sprintf(`create index B on %s."B"("Id","CreatedAt","Age","IdA")`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableC for create new user table in tenant
func (db *Database) CreateTableC(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."C"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableC for insert seed data into table C
func (db *Database) InsertIntoTableC(tenant string) error {
	query := fmt.Sprintf(`insert into %s."C" ("Age") values (%d);`, tenant, 12)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableC for create multi column index in table C
func (db *Database) CreateIndexTableC(tenant string) error {
	query := fmt.Sprintf(`create index C on %s."C"("Id","CreatedAt","Age")`, tenant)

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

// InsertIntoTableD for insert seed data into table D
func (db *Database) InsertIntoTableD(tenant string) error {
	query := fmt.Sprintf(`insert into %s."D" ("Age") values (%d);`, tenant, 12)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableD for create multi column index in table D
func (db *Database) CreateIndexTableD(tenant string) error {
	query := fmt.Sprintf(`create index D on %s."D"("Id","CreatedAt","Age")`, tenant)

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
		"Name" text);`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableE for insert seed data into table E
func (db *Database) InsertIntoTableE(tenant string) error {
	query := fmt.Sprintf(`insert into %s."E" ("Name") values ('%s');`, tenant, "mahdi")
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableE for create multi column index in table E
func (db *Database) CreateIndexTableE(tenant string) error {
	query := fmt.Sprintf(`create index E on %s."E"("Id","CreatedAt","Name")`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableF for create new user table in tenant
func (db *Database) CreateTableF(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."F"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableF for insert seed data into table F
func (db *Database) InsertIntoTableF(tenant string) error {
	query := fmt.Sprintf(`insert into %s."F" ("Age") values (%d);`, tenant, 123)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableF for create multi column index in table F
func (db *Database) CreateIndexTableF(tenant string) error {
	query := fmt.Sprintf(`create index F on %s."F"("Id","CreatedAt","Age")`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableG for create new user table in tenant
func (db *Database) CreateTableG(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."G"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableG for insert seed data into table G
func (db *Database) InsertIntoTableG(tenant string) error {
	query := fmt.Sprintf(`insert into %s."G" ("Age") values (%d);`, tenant, 123)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableG for create multi column index in table G
func (db *Database) CreateIndexTableG(tenant string) error {
	query := fmt.Sprintf(`create index G on %s."G"("Id","CreatedAt","Age")`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableH for create new user table in tenant
func (db *Database) CreateTableH(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."H"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableH for insert seed data into table H
func (db *Database) InsertIntoTableH(tenant string) error {
	query := fmt.Sprintf(`insert into %s."H" ("Age") values (%d);`, tenant, 143)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableH for create multi column index in table H
func (db *Database) CreateIndexTableH(tenant string) error {
	query := fmt.Sprintf(`create index H on %s."H"("Id","CreatedAt","Age")`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableI for create new user table in tenant
func (db *Database) CreateTableI(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."I"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableI for insert seed data into table I
func (db *Database) InsertIntoTableI(tenant string) error {
	query := fmt.Sprintf(`insert into %s."I" ("Age") values (%d);`, tenant, 143)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableI for create multi column index in table I
func (db *Database) CreateIndexTableI(tenant string) error {
	query := fmt.Sprintf(`create index I on %s."I"("Id","CreatedAt","Age")`, tenant)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateTableJ for create new user table in tenant
func (db *Database) CreateTableJ(tenant string) error {
	query := fmt.Sprintf(`create table if not exists %s."J"("Id" serial PRIMARY KEY,
    	"CreatedAt" timestamptz NOT NULL DEFAULT (now()),
		"Age" int);`, tenant)
	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// InsertIntoTableJ for insert seed data into table J
func (db *Database) InsertIntoTableJ(tenant string) error {
	query := fmt.Sprintf(`insert into %s."J" ("Age") values (%d);`, tenant, 143)

	_, err := db.NameSpace.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateIndexTableJ for create multi column index in table J
func (db *Database) CreateIndexTableJ(tenant string) error {
	query := fmt.Sprintf(`create index J on %s."J"("Id","CreatedAt","Age")`, tenant)

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
