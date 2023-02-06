run:
	docker run -d -p 8080:8080 -v data-cities://app/data --env-file ./configs/.env --rm --name cities_service cities_service:v1.0
run-dev:
	docker run -t -i -p 8080:8080 -v data-cities://app/data --env-file ./configs/.env --rm --name cities_service cities_service:v1.0
stop:
	docker stop cities_service
build:
	docker build -t cities_service:v1.0 .
swag:
	swag init -g cmd/app/main.go