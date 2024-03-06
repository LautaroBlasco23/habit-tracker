# This file contains scripts to run and test the app on your machine.


run-docker-database:
	docker run -e POSTGRES_USER=lauti -e POSTGRES_PASSWORD=test123 -p 5432:5432 -d --name=habits-postgres-db postgres 


kill-docker-database:
	docker kill habits-postgres-db


rm-docker-database:
	docker rm habits-postgres-db


run-server-api: 
	HABITS_DB_USER=lauti HABITS_DB_PASSWORD=test123 HABITS_DB_NAME=postgres go run main.go


run-app: run-docker-database run-server-api

