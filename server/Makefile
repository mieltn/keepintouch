CONTAINER_NAME=postgres15
DB_NAME=go-chat
POSTGRES_DSN=postgresql://postgres:secret@localhost:5433/${DB_NAME}

postgres_init:
	docker run -e POSTGRES_PASSWORD=secret -d -p 5433:5432 --name ${CONTAINER_NAME} postgres:15-alpine

createdb:
	docker exec -it ${CONTAINER_NAME} createdb -U postgres ${DB_NAME}

dropdb:
	docker exec -it ${CONTAINER_NAME} dropdb -U postgres ${DB_NAME}

migrate-up:
	goose -dir db/migrations postgres ${POSTGRES_DSN} up

migrate-down:
	goose -dir db/migrations postgres ${POSTGRES_DSN} down

run:
	go build -o build cmd/*.go && ./build/main
