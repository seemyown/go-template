DB_URI ?= postgres://postgres:root@localhost:5432/postgres?sslmode=disable

up:
	migrate -database "$(DB_URI)" -path migrations up

down:
	migrate -database "$(DB_URI)" -path migrations down 1
