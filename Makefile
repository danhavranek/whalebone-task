APP_NAME="whalebone-task"

run:
	go run . 

test:
	go test ./tests

build:
	go build -o bin/${APP_NAME} . 

build-image:
	docker build --tag ${APP_NAME} .

clean:
	rm -fr bin app