.PHONY: go swagger

all: go swagger

go:
	docker-compose run --rm chirpstack-api-go

swagger:
	docker-compose run --rm chirpstack-api-swagger
