start-load: docker-up
	env $(shell cat .env | xargs) go run cmd/main.go load-job

start-api: docker-up
	env $(shell cat .env | xargs)  go run cmd/main.go api

docker-up:
	docker-compose -f docker/docker-compose.yml up -d

docker-down:
	docker-compose -f docker/docker-compose.yml down -v

docker-test-up:
	docker-compose -f docker/docker-compose.test.yml up -d --force-recreate

test: docker-test-up
	env $(shell cat .env.test | xargs) go test -count=1 ./...