
MIGRATION_FOLDER := db/migrations
DB_URL := root:mysql@1234&?@tcp(127.0.0.1:3306)/Airbnb_auth_dev

run:  # make run
	go run main.go

dev:  # make dev
	air
# create a new migration ---> make migrate-create name="create_user_table" ----->>>>>> use gmake on mac
migrate-create:
	goose -dir "$(MIGRATION_FOLDER)" create $(name) sql

migrate-up: # make migrate-up
	goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL}  up

migrate-down: # make migrate-down
	goose -dir ${MIGRATION_FOLDER} mysql ${DB_URL} down

# Rollback all migrations and reset database # make migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" reset

# Show current migration status # gmake migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" status

# Redo last migration (Down then Up) # gmake migrate-redo
migrate-redo:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" redo

# Run specific migration version # gmake migrate-version version=20200101120000
migrate-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version # gmake migrate-down-to version=20200101120000
migrate-down-to:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version # gmake migrate-force version=20200101120000
migrate-force:
	goose -dir $(MIGRATIONS_DIR) mysql "$(DB_URL)" force $(version)

# Print Goose help # gmake migrate-help
migrate-help:
	goose -h

# task1 : write a service layer code, that can encrypt the user password before passing to the repository layer

# task2 : write a service layer function that can create a new JWT token for a user.
