# Carrega variáveis do .env
ifneq (,$(wildcard .env))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

run:
	go run cmd/rest/main.go

test:
	go test ./...

migrate:
	@echo "Running migrations using Docker..."
	@echo "DB_URL=$(DB_URL)"
	@echo "MIGRATIONS_PATH=migrations"
	podman run --rm \
		-v $(CURDIR)/migrations:/migrations:ro,Z \
		--network host \
		migrate/migrate \
		-source=file://migrations \
		-database "$(DB_URL)" \
		up

create-migration:
	@echo "Creating migrations using Docker..."
	@echo "MIGRATIONS_PATH=migrations"
	podman run --rm \
		-v $(CURDIR)/migrations:/migrations \
		--network host \
		migrate/migrate \
		create -ext=sql -dir=migrations \
		-seq init 

gen-private:
	openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048

gen-public:
	openssl rsa -in private.pem -pubout -out public.pem
