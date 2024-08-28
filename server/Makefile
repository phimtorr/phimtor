
.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?parseTime=true" up

.PHONY: sqlboiler
sqlboiler:
	sqlboiler mysql -c sqlboiler.toml

.PHONY: dependencies
dependencies:
	export $(cat .env | grep -v "^#" | xargs) && docker compose up -d
