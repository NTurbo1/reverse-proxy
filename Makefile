EXECUTABLE=reverse_proxy.exe

build:
	@mkdir -p dist && go build -o ./dist/$(EXECUTABLE) ./cmd

run: build
	@./dist/$(EXECUTABLE)

clean:
	@rm -rf ./dist

format:
	@go fmt ./...

.PHONY: run build clean format
