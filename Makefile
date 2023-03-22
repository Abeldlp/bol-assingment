up:
	docker-compose up --build -d
stop:
	docker-compose stop
down:
	docker-compose down
test:
	cd mancala-api && go test ./test/...
	cd api-gateway && go test ./test/...

