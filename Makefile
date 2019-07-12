all: build-vue-app build-go

.PHONY: build-vue-app

build-vue-app:
	npm run --prefix frontend build

build-go:
	packr
	go build -o hello-vue main.go a_main-packr.go