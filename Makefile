CURRENT_DIR := $(shell pwd)
DATABASE_URL := "postgres://postgres:0412@localhost:5432/nt?sslmode=disable"

gen-proto:
	./scripts/gen-proto.sh ${CURRENT_DIR}

tidy:
	@go mod tidy
	@go mod vendor

mig-create:
	@if [ -z "$(name)" ]; then \
		read -p "Enter migration name: " name; \
		migrate create -ext sql -dir migrations -seq $$name; \
	else \
		migrate create -ext sql -dir migrations -seq $(name); \
	fi

mig-up:
	@migrate -database "$(DATABASE_URL)" -path migrations up

mig-down:
	@migrate -database "$(DATABASE_URL)" -path migrations down

mig-force:
	@if [ -z "$(version)" ]; then \
		read -p "Enter migration version: " version; \
		migrate -database "$(DATABASE_URL)" -path migrations force $$version; \
	else \
		migrate -database "$(DATABASE_URL)" -path migrations force $(version); \
	fi
