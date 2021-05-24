run:
	go run main.go

sqlc:
	sqlc generate

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir database/postgres/migration $$name

migrate_up:
	migrate -path database/postgres/migration -database "postgresql://root:secret@localhost:5432/todo_list?sslmode=disable" -verbose up

migrate_down:
	migrate -path database/postgres/migration -database "postgresql://root:secret@localhost:5432/todo_list?sslmode=disable" -verbose down

create_db:
	docker exec -it go_todo_postgres createdb --username=root --owner=root todo_list

drop_db:
	docker exec -it go_todo_postgres dropdb todo_list

test:
	go test -v -cover ./...