build:
	docker compose up --build -d

migrate-up:
	migrate -path migrations -database postgres://postgres:qwerty@localhost:5432/reservation?sslmode=disable up

migrate-down:
	migrate -path migrations -database postgres://postgres:qwerty@localhost:5432/reservation?sslmode=disable down

test:
	cd ./app/internal/service && go test . -v