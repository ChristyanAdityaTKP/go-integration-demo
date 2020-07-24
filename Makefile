build:
	go build .

run:
	go run .

docker-up:
	docker-compose up --build -d
	docker ps

docker-up-test:
	docker-compose up --build -d db-test
	docker ps

docker-down:
	docker-compose down
	docker ps

seed-db:
	chmod +x ./schema/postgres-init.sh
	./schema/postgres-init.sh

integration-test: docker-up-test seed-db
	go test -v ./integration-tests/...

integration-test-ci: docker-up-test seed-db
	go test -v ./integration-tests/...
	docker-compose down
