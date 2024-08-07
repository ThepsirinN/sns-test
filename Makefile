up:
	docker-compose up -d

down:
	docker-compose down

run:
	go run .

tidy:
	go mod tidy

migrate-user:
	go run ./migrate

test-cover:
	go test -cover ./v1/handlers ./v1/services

test-full-cover:
	go test -v -cover ./v1/handlers ./v1/services 