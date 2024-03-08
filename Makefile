# Tools commands

.PHONY: tools
tools:
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
	#pip3 install sabledocs
	#npm install -g @openapitools/openapi-generator-cli



.PHONY: openapi
openapi: openapi-http openapi-client

.PHONY: openapi-http
openapi-http:
	@./scripts/openapi-http.sh phimtor_api server/ports ports

.PHONY: openapi-client
openapi-client:
	@./scripts/openapi-client.sh phimtor_api desktop/client/api api








