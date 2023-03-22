up:
	docker-compose up --build -d
stop:
	docker-compose stop
down:
	docker-compose down
test:
	cd mancala-api && go test ./test/...
	# cd schedule-service && go test ./test
	# cd client && npm run test:unit:ci && npm run test:e2e:ci

