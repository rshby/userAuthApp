migrateup:
	migrate -path db/migration -database "postgres://postgres:root@localhost:5432/userauth_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:root@localhost:5432/userauth_db?sslmode=disable" -verbose down


.PHONY: migrateup migratedown