include .env
export


start:
	@go run main.go

up:
	@docker-compose up -d

down:
	@docker-compose down

sh:
	@docker exec -it `docker ps -q` cqlsh -u ${CASSANDRA_USER} -p ${CASSANDRA_PASSWORD}
