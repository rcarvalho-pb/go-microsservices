FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

## up: start all containers
up:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

## up-build: stops docker (if running), builds all projects and starts docker-compose
up-build: build-broker
	@echo "Stopping docker images (if any running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build-broker: builds the broker binary as linux executable
build-broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## build-front: builds the front end binary
build-front:
	@echo "Building front end binary..."
	cd ./front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## start: starts the front end
start: build-front
	@echo "Starting front end"
	# cd ./front-end && go build -o ${FRONT_END_BINARY} ./cmd/web
	cd ./front-end && ./${FRONT_END_BINARY} &

## stop: stop the front end
stop:
	@echo "Stopping front end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}" &
	@echo "Stopped front end!"
