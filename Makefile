build:
	docker compose up --build -d

migrate-up:
	$(MIGRATE) -path migrations -database postgres://postgres:qwerty@localhost:5432/reservation?sslmode=disable up