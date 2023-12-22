createdb:
	docker exec -t postgres12-budgets createdb --username=root --owner=root budgets
dropdb:
	docker exec -t postgres12-budgets dropdb budgets
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/budgets?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/budgets?sslmode=disable" -verbose down
sqlc:
	cd db && sqlc generate && cd ..
test:
	go test -v -cover ./...
.PHONY: migrateup migratedown createdb dropdb sqlc test