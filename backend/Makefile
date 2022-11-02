run:
	go run ./cmd/main.go

setup:
	go mod download

# Database Migration
# create migration file
migrate-create:
	migrate create -ext sql -dir migration -seq $(NAME)

# migration up (craete all table)
migrate-up:
	migrate -database "mysql://root:fanzru@tcp(103.55.38.98:1000)/skripsi_dev" -path migration up

# migration down (drop all table)
migrate-down:
	migrate -database "mysql://$(MYSQL_DBUSER):$(MYSQL_DBPASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)" -path migration down -all

# rollback migration
migrate-rollback:
	migrate -database "mysql://$(MYSQL_DBUSER):$(MYSQL_DBPASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)" -path migration down $(N)

# migration force with version (craete all table)
migrate-force:
	migrate -database "mysql://$(MYSQL_DBUSER):$(MYSQL_DBPASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)" -path migration force $(VERSION)
