REPOSITORY=jekabolt
REGISTRY=jekabolt
IMAGE_NAME=tolya-robot
VERSION=latest


local:
	grep -lR --exclude=Makefile --exclude-dir=.git  "" . | xargs sed -i 's~http://dotmarket.me~http://localhost:8080~g'

dot:
	grep -lR --exclude=Makefile --exclude-dir=.git  "" . | xargs sed -i 's~http://localhost:8080~http://dotmarket.me~g'

build:
	go build -o ./bin/$(IMAGE_NAME) ./cmd/

run: build
	source .env && ./bin/$(IMAGE_NAME)

image:
	docker build -t $(REPOSITORY)/${IMAGE_NAME}:$(VERSION) -f ./Dockerfile . 
	docker tag $(REPOSITORY)/${IMAGE_NAME}:$(VERSION) $(REGISTRY)/${IMAGE_NAME}:$(VERSION)
