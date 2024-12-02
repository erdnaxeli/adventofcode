all: run

db:
	docker compose up -d

run: db
	#docker compose exec -it db psql -h localhost -U postgres
	psql -h localhost -U postgres
