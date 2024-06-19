postgres-database:
	docker build -t habittracker-postgres-img .
	docker run -d -p 5432:5432 --name habittracker-postgres-db habittracker-postgres-img 

mongo-database:
	docker run --name habittracker-mongodb -p 27017:27017 -d mongo:latest

both-databases-up: mongo-database postgres-database

delete-databases-and-images:
	docker kill habittracker-postgres-db
	docker rm habittracker-postgres-db
	docker rmi habittracker-postgres-img
	docker kill habittracker-mongodb
	docker rm habittracker-mongodb
