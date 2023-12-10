include .env
export

# Artisan

.PHONY: artisan
artisan:
	@go run . artisan $(filter-out $@,$(MAKECMDGOALS))

# OpenAPI Swagger

.PHONY: openapi
openapi:
	@./scripts/openapi-http.sh

# Protobuf

.PHONY: proto
proto:
	@./scripts/proto.sh

# Runner

.PHONY: run
run:
	SERVER_TO_RUN=$(filter-out $@,$(MAKECMDGOALS)) go run .