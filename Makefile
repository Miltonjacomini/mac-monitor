BINARY_NAME=mac-monitor
GUI_BINARY_NAME=mac-monitor-gui
GO_FILES=$(shell find . -name "*.go" -not -path "./vendor/*")

.PHONY: all build gui-build gui-dev test clean run help

all: gui-build

# CLI version build
build:
	CGO_ENABLED=1 go build -o bin/$(BINARY_NAME) main.go

# GUI version build using Wails
gui-build:
	@if command -v wails > /dev/null; then \
		cd gui && wails build -o $(GUI_BINARY_NAME); \
	else \
		echo "Wails CLI not found. Please install it with: go install github.com/wailsapp/wails/v2/cmd/wails@latest"; \
		exit 1; \
	fi

# GUI development mode
gui-dev:
	@cd gui && wails dev

test:
	go test -v -race ./...

clean:
	rm -rf bin/
	rm -rf gui/build/bin/
	go clean

run: build
	./bin/$(BINARY_NAME)

help:
	@echo "Disponíveis:"
	@echo "  make build      - Build da versão CLI"
	@echo "  make gui-build  - Build da versão GUI (Wails)"
	@echo "  make gui-dev    - Inicia desenvolvimento da GUI com Hot Reload"
	@echo "  make test       - Executa testes"
	@echo "  make clean      - Limpa binários"
