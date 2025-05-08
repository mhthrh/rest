IMAGE_NAME=rest_service

buildBinary:
	chmod +x ./script/build.sh
	./script/build.sh cmd/main.go $(IMAGE_NAME)
build: buildBinary
	docker build --progress=plain -t $(IMAGE_NAME) .
run: build
	docker run --rm -p 9090:9090 $(IMAGE_NAME)
