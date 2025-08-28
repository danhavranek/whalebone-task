APP_NAME="whalebone-task"

run:
	go run . 

build:
	go build -o bin/${APP_NAME} . 

build-image:
	docker build --tag ${APP_NAME} .

clean:
	rm -fr bin