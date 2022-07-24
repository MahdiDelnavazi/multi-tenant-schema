postgres:
	docker run --name multi-tenant-image -p 5432:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=secret -d postgres:latest

# create db one for bridge model with separate databases
create-tenant-one-db:
	docker exec -it multi-tenant-image createdb --username=root --owner=root multi-tenant-one

# create db two for bridge model with separate databases
create-tenant-two-db:
	docker exec -it multi-tenant-image createdb --username=root --owner=root multi-tenant-two

# create db schema for bridge model with separate schema
create-tenant-schema-db:
	docker exec -it multi-tenant-image createdb --username=root --owner=root multi-tenant-schema

# drop db schema
drop-tenant-schema-db:
	docker exec -it multi-tenant-image dropdb multi-tenant-schema

# for start server
server:
	go run main.go

# migrate up multi-tenant-schema database
migrateup:
		migrate -path Database/Schema -database "postgresql://root:secret@localhost:5432/multi-tenant-schema?sslmode=disable" -verbose up

.PHONY: postgres create-tenant-one-db create-tenant-two-db create-tenant-schema-db drop-tenant-schema-db server migrateup