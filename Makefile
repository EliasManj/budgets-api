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
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/eliasmanj/budgets-api/db/sqlc Querier
.PHONY: migrateup migratedown createdb dropdb sqlc test mock