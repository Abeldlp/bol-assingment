up:
	if [ ! -d "./mancala-client/node_modules" ]; then cd mancala-client && npm install; fi
	docker-compose up --build -d
stop:
	docker-compose stop
down:
	docker-compose down
test:
	cd mancala-api && go test ./test/...
	cd api-gateway && go test ./test/...
docs:
	cd mancala-api && swag init --parseDependency --parseInternal
