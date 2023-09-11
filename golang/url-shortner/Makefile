DOCKER_IMAGE := "ismael3s/short.io"
DOCKER_IMAGE_TEST := "ismael3s/short.io-test"
APP_NAME := "short.io"

test@unit:
	docker build -t $(DOCKER_IMAGE_TEST) -f Dockerfile.test .
	docker run --rm -it -v $(PWD):/app $(DOCKER_IMAGE_TEST)
build:
	rm -rf $(APP_NAME)
	go mod tidy
	go build -o $(APP_NAME) cmd/main.go
build@docker:
	docker build -t $(DOCKER_IMAGE) .
run@docker:
	docker run --rm -it -p 4000:4000 $(DOCKER_IMAGE)
run@dev:
	go run cmd/main.go
run@prod:
	./$(APP_NAME)