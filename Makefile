IMAGE_NAME=rest_service
Update_File=./script/update.sh
Build_File=./script/build.sh

buildBinary:
	@if [ ! -x "$(Update_File)" ]; then \
    		sudo chmod +x $(Build_File) ;\
    	fi
	./script/build.sh cmd/main.go $(IMAGE_NAME)
build: buildBinary
	docker build --progress=plain -t $(IMAGE_NAME) .
run: build
	docker run --rm -p 9090:9090 $(IMAGE_NAME)

update_lib:
	@if [ ! -x "$(Update_File)" ]; then \
		sudo chmod +x $(Update_File) ;\
	fi
	zsh $(Update_File)
	#sh ./script/update-lib.sh

.PHONY: buildBinary build run update_lib