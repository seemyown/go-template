up:
	migrate -database "postgres://postgres:root@localhost:5432/postgres?sslmode=disable" -path migrations up

down:
	migrate -database "postgres://postgres:root@localhost:5432/postgres?sslmode=disable" -path migrations down 1
