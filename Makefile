APP_NAME=jubelio-chat
BUILD_DIR=./.dist

build:
	CGO_ENABLED go build -ldflags="-w -s" $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)