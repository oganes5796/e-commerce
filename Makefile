migrate:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@0.0.0.0:5436/ecommerce?sslmode=disable' up

migrate_down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@0.0.0.0:5436/ecommerce?sslmode=disable' down

show_migrate:
	psql -h 0.0.0.0 -p 5436 -U postgres -d ecommerce

build:
	docker-compose up --build

run:
	docker-compose up

down:
	docker-compose down

test:
	go test -v ./...