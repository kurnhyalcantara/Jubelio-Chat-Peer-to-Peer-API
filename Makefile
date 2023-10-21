build:
	@go build -o .build/jubelio-chat

run: build
	@./build/jubelio-chat